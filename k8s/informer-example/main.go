package main

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func main() {
	//debugToken := os.Getenv("DEBUG_API_TOKEN")
	//debugAPIHOST := os.Getenv("DEBUG_API_HOST")
	//var cfg *rest.Config
	//if debugToken != "" && debugAPIHOST != "" {
	//	cfg = &rest.Config{
	//		Host: debugAPIHOST,
	//		BearerToken: debugToken,
	//	}
	//}else {
	//	cfgz, err := clientcmd.BuildConfigFromFlags("", "")
	cfg, err := clientcmd.BuildConfigFromFlags("", "/Users/kirago/Documents/CodeRepo/Github/golang-coding/k8s/cfg/config")
	if err != nil {
		panic(err)
	}
	//	cfg = cfgz
	//}
	clientset, err := kubernetes.NewForConfig(cfg)

	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})

	defer close(stopCh)

	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)
	//sharedInformers.Start(stopCh)

	informer := sharedInformers.Core().V1().Pods().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("New Pod Added to Store: %s", mObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			log.Printf("%s Pod updated to %s", oObj.GetName(), nObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Pod Deleted from Store; %s", mObj.GetName())
		},
	})
	informer.Run(stopCh)
}

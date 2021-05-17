package singleton

import "sync"

var (
	instance *Instance
	lock sync.Mutex

	goInstance *Instance
	once sync.Once

)

type Instance struct {
	Name string
}

// 双重检索
func GetInstance(name string) *Instance  {
	if instance == nil{
		lock.Lock()
		defer lock.Unlock()
		if instance == nil{
			instance = &Instance{Name: "test"}
		}
	}
	return instance
}

// go 语言风格的单例模式
func GoInstance(name string) *Instance  {
	if instance == nil{
		// 利用 once.Do 来保证某个对象只初始化一次，哪怕Do函数里面发生异常，对象初始化失败了，这个Do函数也不会再被执行。
		once.Do(func() {
			goInstance = &Instance{Name: "goTest"}
		})
	}
	return goInstance
}


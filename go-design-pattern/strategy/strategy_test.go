package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	o := Operator{}
	o.setStrategy(&add{})
	result := o.calculate(1,2)
	fmt.Println(result)

	o.setStrategy(&reduce{})
	result = o.calculate(2,1)
	fmt.Println(result)
}

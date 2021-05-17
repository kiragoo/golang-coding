package strategy
// 策略模式

// 实现此接口，则为一个策略
type IStrategy interface {
	do(int, int) int
}

// 加
type add struct {}

func (a *add) do(x,y int) int{
	return x+y
}

// 减
type reduce struct {}

func (r *reduce) do(x, y int)int  {
	return x-y
}

// 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (o *Operator) setStrategy(strategy IStrategy)  {
	o.strategy = strategy
}

// 调用策略中的方法
func (o *Operator) calculate(a,b int)  int{
	return o.strategy.do(a,b)
}
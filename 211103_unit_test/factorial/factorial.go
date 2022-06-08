package factorial

// Fact 定义返回的结构体
type Fact struct {
	ret  int
	nums []int
}

// Factorial 计算给定数字的阶乘
func Factorial(x int) *Fact {
	var factS = new(Fact)
	return operation(x, factS)
}

// 递归运算
func operation(x int, f *Fact) *Fact {
	f.nums = append(f.nums, x)
	if x == 1 {
		f.ret = x
		return f
	}
	f.ret = x * operation(x-1, f).ret
	return f
}

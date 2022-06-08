package factorial

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	// 程序输出的结果
	got := Factorial(5)

	// 期望的结果
	want := &Fact{
		ret:  120,
		nums: []int{5, 4, 3, 2, 1},
	}

	// 因为struct不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		// 测试失败输出错误提示
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 这里给到一个错误的结果，测试错误情况
func TestFactEr(t *testing.T) {
	got := Factorial(3)
	want := &Fact{
		ret:  6,
		nums: []int{3, 2, 1},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

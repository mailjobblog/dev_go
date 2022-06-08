package factorial

import (
	"reflect"
	"testing"
)

func TestFactsGroup(t *testing.T) {
	t.Parallel()

	// 定义测试表格
	testCases := []struct {
		testName string
		input    int
		output   *Fact
	}{
		{"case1", 5, &Fact{ret: 120, nums: []int{5, 4, 3, 2, 1}}},
		{"case2", 3, &Fact{ret: 6, nums: []int{3, 2, 1}}},
	}

	setupTest := setupFactsTest(t)
	defer setupTest(t) // 测试之前执行setup操作

	// 运行子测试代码
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			downTest := downFactsTest(t)
			defer downTest(t) // 测试之后执行testdown操作

			got := Factorial(tt.input)
			if !reflect.DeepEqual(tt.output, got) {
				t.Errorf("expected:%v, got:%v", tt.output, got)
			}
		})
	}
}

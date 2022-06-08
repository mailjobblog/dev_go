package factorial

import (
	"testing"
)

// 测试集的Setup与Teardown
func setupFactsTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:Setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:Setup")
	}
}

// 子测试的Setup与Teardown
func downFactsTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:Teardown")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:Teardown")
	}
}

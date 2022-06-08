package factorial

import (
	"reflect"
	"testing"
)

// 工具安装：go get github.com/cweill/gotests/...
// 生成命令：gotests -all -w factorial.go

// gotests 工具生成的测试代码
func Test_operation(t *testing.T) {
	type args struct {
		x     int
		factS *Fact
	}
	tests := []struct {
		name string
		args args
		want *Fact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operation(tt.args.x, tt.args.factS); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

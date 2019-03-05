package stringutil

import "testing" // 导入testing 框架

func TestReverse(t *testing.T) { // 通过对应方法名测试该方法 T 应该为范型
	cases := []struct { // 定义一个结构体 in 代表实参  want string 代表返回类型 且希望为string类型
		in, want string
	}{
		{"Hello ， world", "dlrow ,olleH"}, // 设置三个结构体实例
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases { // foreach 循环 cases
		got := Reverse(c.in) // 定义 一个变量got 且初始为该次执行方法对返回值
		if got != c.want {   // 比对其返回值与测试希望返回值是否一致 如果不对 报错
			t.Error("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestReverse2(t *testing.T) {
	case2 := []struct {
		in, want string
	}{
		{"", ""},
		{"1", "1"},
		{"helloo", "oolleh"},
	}

	for _, c := range case2 {
		result := Reverse(c.in)
		if result != c.want {
			t.Error("Reverse(%q) == %q, want %q", c.in, result, c.want)
		}
	}
}

package name

import (
	"fmt"
	"testing"
)

// 测试用例的四种形式: A, B, C, D

// A: 基本测试
// TestAbc(t *testing.T)

// B: 压力测试
// BenchmarkAbc(b *testing.B)

// C: 测试控制台输出的例子
// ExampleAbc()

// D: 测试Main函数
// TestMain(m *testing.M)

// NamingConversion Naming conversion
type NamingConversion struct {
	Enter string
	Got   string
	Want  string
}

// go test
// go test -v
// go test -run=PascalToUnderline/chinese -v
// 测试覆盖率 => 函数被测试函数执行代码的百分比 如:原函数10行代码,测试函数也执行了10行代码 测试覆盖率为100%,如果测试函数只执行了6行 测试覆盖率则为60%
// go test -cover
// 测试覆盖率结果输出到c.out
// go test -cover -coverprofile=c.out
// html方式查看覆盖率信息
// go tool cover -html=c.out
// go test -cover -coverprofile=c.out && go tool cover -html=c.out
// TestPascalToUnderline 基本测试
func TestPascalToUnderline(t *testing.T) {
	test := map[string]NamingConversion{
		"normal": {
			Enter: "IAmABoy",
			Want:  "i_am_a_boy",
		},
		"unusual": {
			Enter: "_IAmABoy",
			Want:  "_i_am_a_boy",
		},
		"illegal": {
			Enter: "IAMABOY",
			Want:  "i_a_m_a_b_o_y",
		},
		"chinese": {
			Enter: "我是一个男孩",
			Want:  "我是一个男孩",
		},
	}
	for key, val := range test {
		t.Run(key, func(t *testing.T) {
			val.Got = PascalToUnderline(val.Enter)
			if val.Got != val.Want {
				t.Errorf("name:%v failed, want:%v got:%v", key, val.Want, val.Got)
			}
		})
	}
}

// go test -bench=PascalToUnderline
// go test -bench=PascalToUnderline -benchmem
// BenchmarkPascalToUnderline 压力测试
func BenchmarkPascalToUnderline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PascalToUnderline("IAmABoy")
	}
}

// go test -run Example
// ExamplePascalToUnderline 控制台输出测试
func ExamplePascalToUnderline() {
	fmt.Println(PascalToUnderline("IAmABoy"))
	fmt.Println(PascalToUnderline("我是一个男孩"))
	// Output:
	// i_am_a_boy
	// 我是一个男孩
}

// go test
// go test -v
// go test -run=UnderlineToPascal/chinese -v
// go test -cover
// go test -cover -coverprofile=c.out
// go tool cover -html=c.out
// go test -cover -coverprofile=c.out && go tool cover -html=c.out
// TestUnderlineToPascal 基本测试
func TestUnderlineToPascal(t *testing.T) {
	test := map[string]NamingConversion{
		"normal": {
			Enter: "i_am_a_boy",
			Want:  "IAmABoy",
		},
		"unusual": {
			Enter: "_i_am_a_boy",
			Want:  "IAmABoy",
		},
		"unusual1": {
			Enter: "_i_am_a_boy_",
			Want:  "IAmABoy",
		},
		"illegal": {
			Enter: "i_a_m_a_b_o_y",
			Want:  "IAMABOY",
		},
		"chinese": {
			Enter: "我是一个男孩",
			Want:  "我是一个男孩",
		},
	}
	for key, val := range test {
		t.Run(key, func(t *testing.T) {
			val.Got = UnderlineToPascal(val.Enter)
			if val.Got != val.Want {
				t.Errorf("name:%v failed, want:%v got:%v", key, val.Want, val.Got)
			}
		})
	}
}

// go test -bench=UnderlineToPascal
// go test -bench=UnderlineToPascal -benchmem
// BenchmarkUnderlineToPascal 压力测试
func BenchmarkUnderlineToPascal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnderlineToPascal("i_am_a_boy")
	}
}

// go test -run Example
// ExampleUnderlineToPascal 控制台输出测试
func ExampleUnderlineToPascal() {
	fmt.Println(UnderlineToPascal("i_am_a_boy"))
	fmt.Println(UnderlineToPascal("我是一个男孩"))
	// Output:
	// IAmABoy
	// 我是一个男孩
}

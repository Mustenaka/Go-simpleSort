package simplesort

import "testing"

func TestSimplesort(t *testing.T) {
	// 定义测试结构体
	type book struct {
		id    int
		title string
	}

	// 构造一个乱序数组
	books := []book{
		{id: 5, title: "Go语言"},
		{id: 1, title: "Python语言"},
		{id: 2, title: "Ruby语言"},
		{id: 4, title: "Java语言"},
		{id: 3, title: "C语言"},
	}

	Simplesort(books)
}

func TestSimplesortStable(t *testing.T) {
	SimplesortStable([]interface{}{1, 2, 3, 4, 5})
}

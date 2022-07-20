package simplesort

import (
	"fmt"
	"testing"
)

func TestSimplesort(t *testing.T) {
	// define test structure
	type book struct {
		id    int
		title string
	}

	// Construct an out-of-order array
	books := []book{
		{id: 5, title: "Go语言"},
		{id: 1, title: "Python语言"},
		{id: 2, title: "Ruby语言"},
		{id: 4, title: "Java语言"},
		{id: 3, title: "C语言"},
	}

	result, err := Simplesort(books, "id", true)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func TestFunctionTest(t *testing.T) {
	// define test structure
	type book struct {
		id    int
		title string
	}

	// Construct an out-of-order array
	books := []book{
		{id: 5, title: "Go语言"},
		{id: 1, title: "Python语言"},
		{id: 2, title: "Ruby语言"},
		{id: 4, title: "Java语言"},
		{id: 3, title: "C语言"},
	}

	FunctionTest(books)
}

func TestSinTest(t *testing.T) {
	SinTest("title")
}

func TestSimplesortStable(t *testing.T) {
	SimplesortStable([]interface{}{1, 2, 3, 4, 5})
}

package simplesort

import (
	"fmt"
	"reflect"
	"sort"
)

type SortBy []interface{}

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	// 测试一下
	// fmt.Println(reflect.ValueOf(a[i]).Interface())

	// dataType := reflect.TypeOf(a[i])
	dataValue := reflect.ValueOf(a[i])

	for ii := 0; ii < reflect.TypeOf(a[i]).NumField(); ii++ {
		filed := dataValue.Field(ii)
		// filedName := filed.Type().Name()
		// filedValue := dataValue.FieldByName(filedName)
		fmt.Println(filed)
		fmt.Println("-------------------")
	}

	return false
}

// 默认简单排序
func Simplesort(args interface{}) ([]interface{}, error) {
	// 输出数据
	fmt.Println("input: ", args)

	// 将interface{}转换为[]interface{}
	var interfaceSlice []interface{}
	if reflect.TypeOf(args).Kind() == reflect.Slice {
		fmt.Println("args typeof kind: " + reflect.TypeOf(args).Kind().String())
		// fmt.Println(reflect.TypeOf(args).NumField())
		s := reflect.ValueOf(args)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			interfaceSlice = append(interfaceSlice, ele.Interface())
		}
	}

	// 对结构体数组进行排序
	sort.Sort(SortBy(interfaceSlice))

	// 输出数据
	fmt.Println("output: ", interfaceSlice)
	return interfaceSlice, nil
}

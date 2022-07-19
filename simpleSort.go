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
	// dataValue := reflect.ValueOf(a[i])

	for ii := 0; ii < reflect.TypeOf(a[i]).NumField(); ii++ {
		// filed := dataValue.Field(ii)
		// // filedName := filed.Type().Name()
		// // filedValue := dataValue.FieldByName(filedName)
		// fmt.Println(filed)
		// fmt.Println("-------------------")
	}

	return false
}

// @title Simplesort
// @description simplified non-generic sort method.
// @param data interface{} "Data to be sorted."
// @param filedName string "sorted index field name."
// @param order bool "oder or reverse order"
// @return interface{}, error "sort results, nil if no error."
func Simplesort(args interface{}, filedName string, order bool) ([]interface{}, error) {
	// 输出数据
	fmt.Println("input: ", args)

	// interface{} convert to []interface{}
	var interfaceSlice []interface{}

	// 判断类型
	if reflect.TypeOf(args).Kind() != reflect.Slice {
		panic("input need slice kind")
	}

	// 输出类型信息
	fmt.Println("args typeof kind: " + reflect.TypeOf(args).Kind().String())

	// 切片处理
	s := reflect.ValueOf(args)
	for i := 0; i < s.Len(); i++ {
		ele := s.Index(i)
		t := ele.Type()

		interfaceSlice = append(interfaceSlice, ele.Interface())

		for ii := 0; ii < ele.NumField(); ii++ {
			fmt.Printf("name: %s, type: %s, value: %v\n",
				t.Field(ii).Name,
				ele.Field(ii).Type(),
				ele.Field(ii))
			if t.Field(ii).Name == filedName {
				fmt.Println("find filed: " + filedName)
			}
		}
	}

	// 对结构体数组进行排序
	sort.Sort(SortBy(interfaceSlice))

	// 输出数据
	fmt.Println("output: ", interfaceSlice)
	return interfaceSlice, nil
}

// @title SimplesortStable
// @description reflect test
func FunctionTest(data interface{}) {
	getValue := reflect.ValueOf(data) // Value of v
	if getValue.Kind() != reflect.Slice {
		panic("need slice kind")
	}

	l := getValue.Len()
	for i := 0; i < l; i++ {
		value := getValue.Index(i) // Value of item
		typel := value.Type()      // Type of item
		if typel.Kind() != reflect.Struct {
			panic("need struct kind")
		}

		fmt.Printf("type-kind: %s, type-name: %s, value: %v\n", typel.Kind(), typel.Name(), value.Interface())

		num := value.NumField()
		for j := 0; j < num; j++ {
			fmt.Printf("name: %s, type: %s, value: %v\n",
				typel.Field(j).Name,
				value.Field(j).Type(),
				value.Field(j))
		}
	}
}

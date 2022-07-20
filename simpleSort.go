package simplesort

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/Mustenaka/Go-simpleSort/sortConfig"
)

type SortBy []interface{}

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
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

	// 定位目标字段在结构体中的索引
	var index int = 0

	// 切片处理
	s := reflect.ValueOf(args)
	for i := 0; i < s.Len(); i++ {
		ele := s.Index(i)
		t := ele.Type()

		interfaceSlice = append(interfaceSlice, ele.Interface())

		// 检查需要排序字段是否包含在结构体中
		var isExist bool = false
		for ii := 0; ii < ele.NumField(); ii++ {
			// fmt.Printf("name: %s, type: %s, value: %v\n",
			// 	t.Field(ii).Name,
			// 	ele.Field(ii).Type(),
			// 	ele.Field(ii))
			if t.Field(ii).Name == filedName {
				isExist = true
				index = ii
			}
		}

		// 未找到需要排序的字段，抛出错误
		if !isExist {
			panic("not found filed: " + filedName)
		}
	}

	// 处理完成，初始化排序的配置 - 用单例的方式注入Less方法
	sortConfig.CreateInstance(index, filedName, order)

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

// 单例模式测试
func SinTest(filedName string) {
	sortConfig.CreateInstance(2, filedName, true)

	fmt.Println("instance: ", sortConfig.GetInstance().GetIndex())
	fmt.Println("instance: ", sortConfig.GetInstance().GetName())
	fmt.Println("instance: ", sortConfig.GetInstance().GetOrder())
}

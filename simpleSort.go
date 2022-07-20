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
	// Get Sort index
	sortIndex := sortConfig.GetInstance().GetIndex()

	// loaction the index of the target field in the structure
	fieldValue1 := reflect.ValueOf(a[i]).Field(sortIndex)
	fieldValue2 := reflect.ValueOf(a[j]).Field(sortIndex)

	// assert the type of the target field
	switch fieldValue1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fieldValue1.Int() < fieldValue2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fieldValue1.Uint() < fieldValue2.Uint()
	case reflect.Float32, reflect.Float64:
		return fieldValue1.Float() < fieldValue2.Float()
	case reflect.String:
		return fieldValue1.String() < fieldValue2.String()
	// not support other type
	default:
		panic("unsupported kind")
	}
}

// @title Simplesort
// @description simplified non-generic sort method.
// @param data interface{} "Data to be sorted."
// @param filedName string "sorted index field name."
// @param order bool "oder or reverse order"
// @return interface{}, error "sort results, nil if no error."
func Simplesort(args interface{}, filedName string, order bool) ([]interface{}, error) {
	// print input data(args)
	// fmt.Println("input: ", args)

	// interface{} convert to []interface{}
	var interfaceSlice []interface{}

	// assest type of args, it must be struct. (feat) need support other type.
	if reflect.TypeOf(args).Kind() != reflect.Slice {
		// panic("input need slice kind")
		return nil, fmt.Errorf("input need slice kind")
	}

	// locate the index of the target field in the structure
	var index int = 0

	// get the value of the target field, and convert it to []interface{}
	s := reflect.ValueOf(args)
	for i := 0; i < s.Len(); i++ {
		ele := s.Index(i)
		t := ele.Type()

		interfaceSlice = append(interfaceSlice, ele.Interface())

		// check the type of the target field
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

		// check the target field is not exist
		if !isExist {
			// panic("not found filed: " + filedName)
			return nil, fmt.Errorf("not found filed: %s", filedName)
		}
	}

	// Initialize the configuration of sorting - inject the Less() with a singleton
	sortConfig.CreateInstance(index, filedName, order)

	// sort
	sort.Sort(SortBy(interfaceSlice))

	// output and return the result
	// fmt.Println("output: ", interfaceSlice)
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

// test instance of Simplesort
func SinTest(filedName string) {
	sortConfig.CreateInstance(2, filedName, true)

	fmt.Println("instance: ", sortConfig.GetInstance().GetIndex())
	fmt.Println("instance: ", sortConfig.GetInstance().GetName())
	fmt.Println("instance: ", sortConfig.GetInstance().GetOrder())
}

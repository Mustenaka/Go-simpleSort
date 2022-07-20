package sortConfig

type sortConfig struct {
	index int    // 字段索引
	name  string // 字段名称
	order bool   // 排序方式 - true顺序，false逆序
}

type SortConfig interface {
	// 获取index
	GetIndex() int
	// 获取name
	GetName() string
}

var _ SortConfig = (*sortConfig)(nil)

var instance *sortConfig

// 获取单例
func CreateInstance(index int, name string, order bool) *sortConfig {
	if instance == nil {
		instance = NewSortConfig(index, name, order)
	}
	return instance
}

// 获取单例
func GetInstance() *sortConfig {
	return instance
}

// 初始化
func NewSortConfig(index int, name string, order bool) *sortConfig {
	return &sortConfig{
		index: index,
		name:  name,
		order: order,
	}
}

func (conf *sortConfig) GetIndex() int {
	return conf.index
}

func (conf *sortConfig) GetName() string {
	return conf.name
}

func (conf *sortConfig) GetOrder() bool {
	return conf.order
}

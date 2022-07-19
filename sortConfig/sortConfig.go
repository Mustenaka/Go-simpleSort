package sortconfig

type sortConfig struct {
	index int    // 字段索引
	name  string // 字段名称
}

type SortConfig interface {
	// 获取index
	GetIndex() int
	// 获取name
	GetName() string
}

var _ SortConfig = (*sortConfig)(nil)

// 初始化
func NewSortConfig(index int, name string) *sortConfig {
	return &sortConfig{
		index: index,
		name:  name,
	}
}

func (conf *sortConfig) GetIndex() int {
	return conf.index
}

func (conf *sortConfig) GetName() string {
	return conf.name
}

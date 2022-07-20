package sortConfig

type sortConfig struct {
	index int
	name  string
	order bool // true: ascending, false: descending
}

type SortConfig interface {
	GetIndex() int
	GetName() string
	GetOrder() bool
}

var _ SortConfig = (*sortConfig)(nil)
var instance *sortConfig

// CreateSortConfig instance
func CreateInstance(index int, name string, order bool) *sortConfig {
	if instance == nil {
		instance = NewSortConfig(index, name, order)
	}
	return instance
}

// GetSortConfig instance
func GetInstance() *sortConfig {
	return instance
}

// initlize the configuration of sorting
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

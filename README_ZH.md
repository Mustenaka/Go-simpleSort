# Go-simpleSort（Go简单排序）

Golang在1.18版本之前的排序由于不支持泛型，因此传入参数相当繁杂，还需要自己定义less，swap等方法，本包用于快速解决这一问题，实现一个方法的快速排序。

设计思路笔记：
1. 传入使用interface{}
2. interface{}转换为golang sort支持的[]interface{}做一个鸡肋的泛型效果
3. 每一个具体的值通过field取出来，然后比较，执行sortBy
4. 转换interface{} 然后返回


simplesort传入的参数类型：
1. interface{}  数据本身
2. string     字符串类型，表示需要比较的字段名称（自动识别类型）
3. bool         布尔类型，true表示升序，false表示降序

```
支持的自动比较字段类型：
int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool，time.Time
```

后续 2. string修改为interface{}可以传递多组字段，表示排序的前后关系
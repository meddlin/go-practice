package DynamicArray

/*
*
* Notes:
* - creating a strcut: https://golangbot.com/structs-instead-of-classes/
 */
type DynamicArray struct {
	// capacity int
	// length   int
	data []int
}

/*
* Will initialize an empty array with a capacity of `capacity`, capacity > 0
*
* Notes:
* - making constructors: https://asankov.dev/blog/2022/01/29/different-ways-to-initialize-go-structs/#constructor-function
 */
func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{data: make([]int, capacity)}
}

/*
* Returns the element at index i. (Assuming index i is valid)
 */
func (da DynamicArray) get(i int) int {
	return da.data[i]
}

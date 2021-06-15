package emath

// DeleteUint16ByValue delete the first item from uint16 slice by value, then return the new slice.
func DeleteUint16ByValue(arr []uint16, value uint16) []uint16 {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return DeleteUint16ByIndex(arr, i)
		}
	}
	return arr
}

// DeleteUint16ByIndex Delete value from uint16 by its index, then return the new slice.
func DeleteUint16ByIndex(arr []uint16, index int) []uint16 {
	if index < len(arr)-1 {
		copy(arr[index:], arr[index+1:])
	}
	return arr[:len(arr)-1]
}

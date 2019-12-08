package utils

// MinInt find and return the minimum item from the specified int
func MinInt(nums ...int) int {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxInt find and return the maximum item from the specified int
func MaxInt(nums ...int) int {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinByte find and return the minimum item from the specified byte
func MinByte(nums ...byte) byte {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxByte find and return the maximum item from the specified byte
func MaxByte(nums ...byte) byte {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinInt8 find and return the minimum item from the specified int8
func MinInt8(nums ...int8) int8 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxInt8 find and return the maximum item from the specified int8
func MaxInt8(nums ...int8) int8 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinUint8 find and return the minimum item from the specified uint8
func MinUint8(nums ...uint8) uint8 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxUint8 find and return the maximum item from the specified uint8
func MaxUint8(nums ...uint8) uint8 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinInt16 find and return the minimum item from the specified int16
func MinInt16(nums ...int16) int16 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxInt16 find and return the maximum item from the specified int16
func MaxInt16(nums ...int16) int16 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinUint16 find and return the minimum item from the specified uint16
func MinUint16(nums ...uint16) uint16 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxUint16 find and return the maximum item from the specified uint16
func MaxUint16(nums ...uint16) uint16 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinInt32 find and return the minimum item from the specified int32
func MinInt32(nums ...int32) int32 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxInt32 find and return the maximum item from the specified int32
func MaxInt32(nums ...int32) int32 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinUint32 find and return the minimum item from the specified int32
func MinUint32(nums ...uint32) uint32 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxUint32 find and return the maximum item from the specified uint32
func MaxUint32(nums ...uint32) uint32 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinInt64 find and return the minimum item from the specified uint32
func MinInt64(nums ...int64) int64 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxInt64 find and return the maximum item from the specified int64
func MaxInt64(nums ...int64) int64 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinUint64 find and return the minimum item from the specified uint64
func MinUint64(nums ...uint64) uint64 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxUint64 find and return the maximum item from the specified uint64
func MaxUint64(nums ...uint64) uint64 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinFloat32 find and return the minimum item from the specified float32
func MinFloat32(nums ...float32) float32 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxFloat32 find and return the maximum item from the specified float32
func MaxFloat32(nums ...float32) float32 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MinFloat64 find and return the minimum item from the specified float64
func MinFloat64(nums ...float64) float64 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

// MaxFloat64 find and return the maximum item from the specified float64
func MaxFloat64(nums ...float64) float64 {
	var ret = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

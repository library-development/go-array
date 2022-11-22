package array

func Map[FromType, ToType any](arr []FromType, f func(from FromType) ToType) []ToType {
	result := make([]ToType, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

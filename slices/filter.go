package slices

func Filter[I any](inp []I, filterFunc func(int, I) bool) []I {
	res := make([]I, 0, len(inp))
	for i := range inp {
		if filterFunc(i, inp[i]) {
			res = append(res, inp[i])
		}
	}
	return res
}

// Типовые вспомогательные функции для фильтрации срезов

func NotEmptyStringFilterFunc(_ int, s string) bool {
	return s != ""
}

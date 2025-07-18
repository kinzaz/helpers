package values

type conditional[R any] struct {
	trueVal   R
	condition bool
}

func Assign[R any](trueVal R) *conditional[R] {
	return &conditional[R]{trueVal: trueVal}
}
func (o *conditional[R]) If(condition bool) *conditional[R] {
	o.condition = condition
	return o
}
func (o *conditional[R]) Else(falseVal R) R {
	if o.condition {
		return o.trueVal
	}
	return falseVal
}

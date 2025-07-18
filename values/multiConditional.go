package values

type multiConditional[R any] struct {
	condValues []R
	conditions []bool
}

func AssignOne[R any]() *multiConditional[R] {
	return new(multiConditional[R])
}
func (o *multiConditional[R]) If(condition bool) *multiConditional[R] {
	o.conditions = append(o.conditions, condition)
	return o
}
func (o *multiConditional[R]) Then(trueValue R) *multiConditional[R] {
	o.condValues = append(o.condValues, trueValue)
	return o
}
func (o *multiConditional[R]) Else(falseValue R) R {
	for i := range o.conditions {
		if o.conditions[i] {
			return o.condValues[i]
		}
	}
	return falseValue
}

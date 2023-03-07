package slices

type (
	Ints []int
)

func FromIntSlice(s []int) Ints {
	return Ints(s)
}

func (i Ints) Map(fn func(value, index int) int) Ints {
	ints := make(Ints, len(i))
	for k, v := range i {
		ints[k] = fn(v, k)
	}

	return ints
}

func (i Ints) Filter(fn func(value, index int) bool) Ints {
	var ints Ints
	for k, v := range i {
		if fn(v, k) {
			ints = append(ints, v)
		}
	}

	return ints
}

func (i Ints) ToIntSlice() []int {
	return []int(i)
}

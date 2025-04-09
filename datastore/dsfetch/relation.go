package dsfetch

// RelationList is a list wrapper for ValueCollection adding
// direct access to its values
type RelationList[C any, T loader[C]] struct {
	refs []*ValueCollection[C, T]
}

func (r *RelationList[T, P]) Refs() []*ValueCollection[T, P] {
	return r.refs
}

func (r *RelationList[T, P]) Values() []P {
	values := make([]P, len(r.refs))
	for i, ref := range r.refs {
		values[i] = ref.Get()
	}

	return values
}

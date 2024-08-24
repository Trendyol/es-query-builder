package condition

func If[T ~map[string]any](item T, condition bool) T {
	if !condition {
		return nil
	}
	return item
}

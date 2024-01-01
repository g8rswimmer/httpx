package field

func Set[M ~map[string]V, V any](obj M) map[string]struct{} {
	set := map[string]struct{}{}
	for k := range obj {
		set[k] = struct{}{}
	}
	return set
}

package panic

func parent(fn func()) (recovered bool) {
	defer func() {
		if r := recover(); r != nil {
			recovered = true
		}
	}()
	fn()
	return recovered
}

func panicChild() {
	panic("child problem")
}

func child() {
}

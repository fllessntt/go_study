package study

func TestSwap() {
	a, b := 1, 2
	a, b = b, a
	println(a, b)
	swap(&a, &b)
	println(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

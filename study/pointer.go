package study

func TestPointer() {
	var i = 1
	var j = &i
	var k = j

	println(i, *j, *k)
}

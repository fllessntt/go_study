package study

type person struct {
	name   string
	age    uint8
	gender uint8 //0表示女, 1表示男
}

func TestStruct() {
	person1 := person{"周润发", 68, 1}
	person2 := person{age: 65, name: "丁勇岱", gender: 1}
	println(person1.name, person2.name)

	var personPointer *person
	personPointer = &person2
	println(personPointer.name)
}

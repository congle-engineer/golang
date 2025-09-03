package main
import ("fmt")

const PI = 3.14

const (
	A int = 1
	B = 3.14
	C = "Hi"
)

type Person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	fmt.Println("Hello World!")

	var student1 string = "Cong"
	var student2 = "Phuong"
	x := 2

	fmt.Println("student1: ", student1)
	fmt.Println("student2: ", student2)
	fmt.Println("x: ", x)

	var a string
	var b int
	var c bool

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)

	student1 = "John"
	fmt.Println("student1: ", student1)

	var d, e, f, g int = 1, 3, 5, 7
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)

	var h, i = 10, "Hello"
	fmt.Println("h: ", h)
	fmt.Println("i: ", i)

	j, k := 20, "World"
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)

	var (
		l int
		m int = 1
		n string = "Test"
	)
	fmt.Println("l: ", l)
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)

	var p, q string = "Hello", "World"
	fmt.Print(p)
	fmt.Print(q)

	fmt.Print("\n")

	fmt.Print(p, "\n")
	fmt.Print(q, "\n")

	fmt.Print(p, "\n", q, "\n")

	fmt.Print(p, " ", q, "\n")

	var r, s = 10, 20
	fmt.Print(r, s, "\n")

	fmt.Println(p, q)

	var y string = "Hello"
	var z int = 15

	fmt.Printf("y has value: %v and type: %T\n", y, y)
	fmt.Printf("z has value: %v and type: %T\n", z, z)

	fmt.Printf("%b\n", z)
	fmt.Printf("%d\n", z)
	fmt.Printf("%+d\n", z)
	fmt.Printf("%o\n", z)
	fmt.Printf("%O\n", z)
	fmt.Printf("%x\n", z)
	fmt.Printf("%X\n", z)
	fmt.Printf("%#x\n", z)
	fmt.Printf("%4d\n", z)
	fmt.Printf("%-4d\n", z)
	fmt.Printf("%04d\n", z)

	var txt = "Hello"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	var testBool = true
	fmt.Printf("%t\n", testBool)

	var testFloat = 3.141
	fmt.Printf("%e\n", testFloat)
	fmt.Printf("%f\n", testFloat)
	fmt.Printf("%.2f\n", testFloat)
	fmt.Printf("%6.2f\n", testFloat)
	fmt.Printf("%g\n", testFloat)

	var arr1 = [6]int{10, 11, 12, 13, 14, 15}
	arr2 := [5]int{4, 5, 6, 7, 8}

	arr3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr3[0])

	arr4 := [5]int{1:10, 2:100}
	fmt.Println(arr4)
	fmt.Println(len(arr4))

	myslice1 := []int{}
	fmt.Println("len myslice1: ", len(myslice1))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println("len myslice2: ", len(myslice2))

	myslice := arr1[2:4]
	fmt.Println("myslice: ", myslice)
	fmt.Println("len myslice: ", len(myslice))
	fmt.Println("cap myslice: ", cap(myslice))

	myslice3 := make([]int, 5, 10)
	fmt.Println("myslice3: ", myslice3)
	fmt.Println("len myslice3: ", len(myslice3))
	fmt.Println("cap myslice3: ", cap(myslice3))

	myslice = append(myslice, 20, 21)
	fmt.Println("myslice: ", myslice)

	mySliceA := []int{1, 2, 3}
	mySliceB := []int{4, 5, 6}
	mySliceC := append(mySliceA, mySliceB...)
	fmt.Println("mySliceC: ", mySliceC)

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

	neededNumbers := numbers[:len(numbers) - 10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Println("numbersCopy: ", numbersCopy)

	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	time := 22
	if (time < 18) {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	myMessage()

	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")

	fmt.Println(sum(10, 20))
	fmt.Println(multipleOutput())

	var pers1 Person
	var pers2 Person

	pers1.name = "John"
	pers1.age = 25
	pers1.job = "Software Engineer"
	pers1.salary = 5000

	pers2.name = "Jane"
	pers2.age = 30
	pers2.job = "Software Engineer"
	pers2.salary = 6000

	fmt.Println("pers1: ", pers1)
	fmt.Println("pers2: ", pers2)

	var map1 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	fmt.Println(map1)
}

func myMessage() {
	fmt.Println("I just got executed!")
}

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func sum(x int, y int) int {
	return x + y
}

func multipleOutput() (int, string) {
	x := 10
	y := "Hello"
	return x, y
}
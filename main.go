package main
import ("fmt")

const PI = 3.14

const (
	A int = 1
	B = 3.14
	C = "Hi"
)

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
}

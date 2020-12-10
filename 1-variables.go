// software entry point "main"
// software intended to be an ancillary lib will not require a main
// external libs will be executed from within a programs main (somewhere in the execution)
package main

// std-lib I/O
// singular import
import "fmt"

// inferred typing
var (
	one   = 1
	two   = 2
	three = 3
)

// declared type - uninitialized
var (
	hello string
	world string
)

// multiline inferred typing
var foo = "foo"
var bar = "bar"

// shorthand inferred typing
var baz, rab = "baz", "rab"

// initialized with typing
var oof string = "oof"

func main() {
	fmt.Println(one, two, three)
	fmt.Println(hello, world)
	fmt.Println(foo, bar)
	fmt.Println(baz, rab)
	fmt.Println(oof)

	// shorthand inferred typing declaration
	// := construct is only available from within a function declaration
	some, value := "some", 20
	fmt.Println(some, value)
}

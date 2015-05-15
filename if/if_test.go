package ifff

import "testing"

//. "github.com/rshmelev/go-ternary/if"

func TestAverage(t *testing.T) {

	println("should be empty, not wow: ", If(0 > 1).Then("wow").Str())

	println("should be 1: ", Or(1).Or(2).Str())
	println("should be 2: ", Or(0).Or(2).Str())
	println("should be hi: ", Or("").Or("hi").Str())
	println("should be hi: ", Or("").Or("hi").Or("bye").Str())
	println("should be hi: ", Or("hi").Or("hey").Or("bye").Str())

	a := 0
	IfDefaultValue(&a).ThenSetTo(10)
	println(a)
	println("------------------")
	println(If(3 > 2).Then("11111").Else(222222).Str())
	println(If(3 > 2).Then("11111").Else(222222).Int())
	println(If(1 > 2).Then("11111").Else(222222).Str())
	println(If(1 > 2).Then("11111").Else(222222).Int())

	println(If(3 > 2).Then("11111").Else(222222).Float32())
	println(If(3 > 2).ThenElse("11111", 222222).Float32())
}

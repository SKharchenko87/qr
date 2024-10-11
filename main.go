package main

func main() {
	println(0b0100)
	println(0b0100_0000)
	println(0b1000_0000)
	println(^0b1111_1111)
	println("===")
	x := uint8(0b11101100)
	println(x)
	println(^x)
	println("===")
	x = uint8(0b00010001)
	println(x)
	println(^x)
}

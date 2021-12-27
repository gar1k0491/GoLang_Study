package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	/*
	   ======================================================================
	   	var a int8               // -128 to 127 (1 byte or 8 bit in memory)
	   	var b int16              // -32768 to 32767 (2 bytes or 16 bit in memory)
	   	var c int32              // -2bil to 2bil (4 bytes or 32 bit in memory)
	   	var d int64              // -9pent to 9pent (8 bytes or 64 bit in memory)
	   //-----------------------
	   	var e uint8              // 0 to 255 (1 byte or 8 bit in memory)
	   	var f uint16             // 0 to 65535 (2 bytes or 16 bit in memory)
	   	var g uint32             // 0 to 4bil (4 bytes or 32 bit in memory)
	   	var h uint64             // 0 to 18pent (8 bytes or 64 bit in memory)
	   //-----------------------
	   	var i byte               // synonym uint8
	   	var j rune               // synonym int32
	   	var k int                // if u have win86 ur int = int32 and if u have win64 ur int = uint64
	   	var l uint               // if u have win86 ur uint = uint32 and if u have win64 ur uint = uint64
	   //-----------------------
	   	var m float32            // 1.4*10^-45 to 3.4*10^35 (4 bytes or 32 bit in memory)
	   	var n float64            // 4.8*10^-320 to 1.8*10^305 (8 bytes or 64 bit in memory)
	   	var o complex64 = 1+2i   //
	   	var p complex128 = 4+90i //
	   //-----------------------
	   	var q bool = true
	   	var r bool = false
	   //-----------------------
	   	var name string = "text"
	   ======================================================================
	*/
	var (
		name = "Igor"
		age  = 32
	)
	var c = fmt.Sprintf("My name is %s and i am %d years old.", name, age)
	fmt.Println(c)

}

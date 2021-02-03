package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	dst := make([]byte, 10)
	src := make([]byte, 10)

	// write with direct assignment
	src[0] = byte(uint8(1))
	src[1] = byte(uint8(5))
	src[2] = byte(uint8(9))

	// write with copying
	buf32 := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf32, uint32(1000000))
	copy(src[3:], buf32)

	fmt.Printf("src: %v\n", src)
	// copy all the bytes
	copy(dst, src)
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)

	buf := bytes.NewBuffer(dst)

	int1, err := binary.ReadUvarint(buf)
	fmt.Println("first one: ", int1, err)

	int2, err := binary.ReadUvarint(buf)
	fmt.Println("second one: ", int2, err)

	int3, err := binary.ReadUvarint(buf)
	fmt.Println("third one: ", int3, err)

	var int4 uint32
	err = binary.Read(buf, binary.LittleEndian, &int4)
	fmt.Println("fourth one: ", int4, err)
}

// from encoding/binary
// func (littleEndian) PutUint32(b []byte, v uint32) {
// 	_ = b[3] // early bounds check to guarantee safety of writes below
// 	b[0] = byte(v)
// 	b[1] = byte(v >> 8)
// 	b[2] = byte(v >> 16)
// 	b[3] = byte(v >> 24)
// }

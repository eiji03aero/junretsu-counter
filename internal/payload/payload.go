package payload

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

type Payload struct {
	ButtonId uint8
}

var (
	payloadSize = reflect.TypeOf(Payload{}).Size()
)

func Encode(buttonId uint8) []byte {
	buf := make([]byte, payloadSize)
	buf[0] = byte(buttonId)
	return buf
}

func Decode(payl *Payload, src []byte) error {
	buffer := bytes.NewBuffer(src)
	err := binary.Read(buffer, binary.LittleEndian, &(payl.ButtonId))
	return err
}

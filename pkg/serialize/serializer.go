package serialize

import (
	"fmt"
	"unsafe"
)

// Marshal used for marshaling uint64 array to []byte
//
//	func Marshal(src []uint64) (data []byte, err error) {
//		buff := &bytes.Buffer{}
//		writer := bufio.NewWriter(buff)
//		temp := make([]byte, 8)
//		for _, i := range src {
//			binary.BigEndian.PutUint64(temp, i)
//			_, err := writer.Write(temp)
//			if err != nil {
//				return nil, err
//			}
//		}
//		_ = writer.Flush()
//		return buff.Bytes(), nil
//	}
func Marshal(src []uint64) (data []byte, err error) {
	arraySize := len(src) * 8
	d := make([]byte, arraySize)
	index := 0
	for _, item := range src {
		for idx := 0; idx < int(unsafe.Sizeof(item)); idx++ {
			byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&item)) + uintptr(idx)))
			d[index] = byt
			index++
		}
	}

	return d, nil
}

// Unmarshal used for unmarshaling [] byte to uint64 array
func Unmarshal(data []byte) ([]uint64, error) {
	l := len(data)
	if l%8 != 0 {
		return nil, fmt.Errorf("number of bytes must be odd")
	}
	ia := make([]uint64, l/8)
	//dst = &[]uint64{}
	for idx := 0; idx < l/8; idx++ {
		val := uint64(0)
		for p := 0; p < 8; p++ {
			*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(p))) = data[idx*8+p]
		}
		ia[idx] = val
	}
	return ia, nil
}

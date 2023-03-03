package pcodec

import (
	"github.com/Gishe8-Team/proto/pkg/serialize"
	"github.com/tilinna/z85"
)

func Encode(src []uint64) (string, error) {
	data, err := serialize.Marshal(src)
	if err != nil {
		return "", err
	}
	buf := make([]byte, z85.EncodedLen(len(src)*8))
	_, eer := z85.Encode(buf, data)
	if eer != nil {
		return "", eer
	}
	return string(buf), nil
}

func Decode(src string) ([]uint64, error) {
	data := make([]byte, z85.DecodedLen(len(src)))
	_, err := z85.Decode(data, []byte(src))
	if err != nil {
		return nil, err
	}

	ia, uer := serialize.Unmarshal(data)
	if uer != nil {
		return nil, uer
	}
	return ia, nil
}

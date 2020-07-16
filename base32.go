package clockwork

import (
	"bytes"
	"fmt"
	"github.com/bearmini/bitstream-go"
)

// TODO: hyphen

func Encode(data []byte) []byte {
	len := len(data)
	r := bitstream.NewReader(bytes.NewReader(data), nil)

	var buf bytes.Buffer
	i := len * 8
	for ; i >= 5; i -= 5 {
		b, err := r.ReadNBitsAsUint8(5)
		if err != nil {
			break
		}
		buf.WriteByte(encodeSymbols[b])
	}
	if i > 0 {
		b, err := r.ReadNBitsAsUint8(uint8(i))
		if err != nil {
		} else {
			b <<= uint(5 - i)
			buf.WriteByte(encodeSymbols[b])
		}
	}
	return buf.Bytes()
}

func Decode(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := bitstream.NewWriter(&buf)

	dataLen := len(data)
	padding := dataLen * 5 % 8
	n := 5
	total := 0
	for i := 0; i < dataLen; i++ {
		b := data[i]
		sym := decodeSymbols[b]
		if sym < 0 {
			return nil, fmt.Errorf("invalid symbol value %c", sym)
		}
		if i+1 == dataLen && padding > 0 {
			n = 5 - padding
			sym >>= uint8(padding)
		}
		w.WriteNBitsOfUint8(uint8(n), uint8(sym))
		total += n
	}

	if total%8 > 0 {
		return nil, fmt.Errorf("invalid total number of decoded bits %d", total)
	}
	return buf.Bytes(), nil
}

var encodeSymbols = [32]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K',
	'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X',
	'Y', 'Z',
}

var decodeSymbols = [256]int8{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 0-9 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 10-19 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 20-29 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 30-39 */
	-1, -1, -1, -1, -1, -1, -1, -1, 0, 1, /* 40-49 */
	2, 3, 4, 5, 6, 7, 8, 9, 0, -1, /* 50-59 */
	-1, -1, -1, -1, -1, 10, 11, 12, 13, 14, /* 60-69 */
	15, 16, 17, 1, 18, 19, 1, 20, 21, 0, /* 70-79 */
	22, 23, 24, 25, 26, -2, 27, 28, 29, 30, /* 80-89 */
	31, -1, -1, -1, -1, -1, -1, 10, 11, 12, /* 90-99 */
	13, 14, 15, 16, 17, 1, 18, 19, 1, 20, /* 100-109 */
	21, 0, 22, 23, 24, 25, 26, -1, 27, 28, /* 110-119 */
	29, 30, 31, -1, -1, -1, -1, -1, -1, -1, /* 120-129 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 130-109 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 140-109 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 150-109 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 160-109 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 170-109 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 180-109 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 190-109 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 200-209 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 210-209 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 220-209 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 230-209 */
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, /* 240-209 */
	-1, -1, -1, -1, -1, -1, /* 250-256 */
}

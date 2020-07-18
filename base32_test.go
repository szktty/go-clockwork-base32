package clockwork

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, testCase := range testCases {
		encoded := Encode(testCase.plain)
		if bytes.Compare(encoded, testCase.encoded) != 0 {
			t.Errorf("encoded '%s', expected '%s', actual '%s'\n",
				testCase.plain, testCase.encoded, encoded)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, testCase := range testCases {
		decoded, err := Decode(testCase.encoded)
		if err != nil {
			t.Errorf("error => %s", err)
		}
		if bytes.Compare(decoded, testCase.plain) != 0 {
			t.Errorf("decoded '%s', expected '%s', actual '%s'\n",
				testCase.encoded, testCase.plain, decoded)
		}
	}
}

type testCase struct {
	plain []byte
	encoded []byte
}

var testCases = []testCase {
	{ []byte("foobar"), []byte("CSQPYRK1E8") },
	{ []byte("Hello, world!"),
		[]byte("91JPRV3F5GG7EVVJDHJ22") },
	{ []byte("The quick brown fox jumps over the lazy dog."),
		[]byte("AHM6A83HENMP6TS0C9S6YXVE41K6YY10D9TPTW3K41QQCSBJ41T6GS90DHGQMY90CHQPEBG") },
	{ []byte("Wow, it really works!"),
		[]byte("AXQQEB10D5T20WK5C5P6RY90EXQQ4TVK44") },
}

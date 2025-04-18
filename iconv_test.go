// iconv_test.go
package iconv

import (
	"fmt"
	"testing"
)

type testCase struct {
	utf8 string

	encodingName string // encodingName that libiconv accepts
	encoded      string // hex representation of encoded string in encodingName

	oneway bool // true if roundtrip conversion doesn't work
}

func (tc testCase) String() string {
	return fmt.Sprintf("%s(%s)", tc.utf8, tc.encodingName)
}

var testData = []testCase{
	{
		utf8:         "これは漢字です。",
		encodingName: "SJIS",
		encoded:      "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B",
	},
	{
		utf8:         "これは漢字です。",
		encodingName: "UTF-16LE",
		encoded:      "S0\x8c0o0\"oW[g0Y0\x020",
	},
	{
		utf8:         "これは漢字です。",
		encodingName: "UTF-16BE",
		encoded:      "0S0\x8c0oo\"[W0g0Y0\x02",
	},
	{
		utf8:         "€1 is cheap",
		encodingName: "ISO-8859-15",
		encoded:      "\xa41 is cheap",
	},
	{
		utf8:         "猫",
		encodingName: "SJIS",
		encoded:      "\x94\x4c",
	},
	{
		utf8:         "①②",
		encodingName: "EUC-JISX0213", // this is only available when GNU libiconv is built with  --enable-extra-encodings
		encoded:      "\xad\xa1\xad\xa2",
	},
	{
		utf8:         "①②",
		encodingName: "EUC-JIS-2004", // this is only available when GNU libiconv is built with  --enable-extra-encodings
		encoded:      "\xad\xa1\xad\xa2",
	},
	{
		utf8:         "髙ⅰ",
		encodingName: "ISO-2022-JP-MS", // this is only available in GNU libiconv
		encoded:      "\x1b$B|b|q\x1b(B",
		oneway:       true,
	},
}

func TestIconv(t *testing.T) {
	for _, data := range testData {
		t.Run(data.String(), func(t *testing.T) {
			cd, err := Open("UTF-8", data.encodingName)
			if err != nil {
				t.Errorf("Error on opening: %s (perhaps you're not using GNU libiconv?)", err)
				return
			}

			str, err := cd.Conv(data.encoded)
			if err != nil {
				t.Errorf("Error on conversion: %s", err)
				return
			}

			if str != data.utf8 {
				t.Errorf("Unexpected value: %#v (expected %#v)", str, data.utf8)
				return
			}

			err = cd.Close()
			if err != nil {
				t.Errorf("Error on close: %s", err)
			}
		})
	}
}

func TestIconvReverse(t *testing.T) {
	for _, data := range testData {
		t.Run(data.String(), func(t *testing.T) {
			if data.oneway {
				t.Skip("roundtrip conversion doesn't work")
			}

			cd, err := Open(data.encodingName, "UTF-8")
			if err != nil {
				t.Errorf("Error on opening: %s (perhaps you're not using GNU libiconv?)", err)
				return
			}

			str, err := cd.Conv(data.utf8)
			if err != nil {
				t.Errorf("Error on conversion: %s", err)
				return
			}

			if str != data.encoded {
				t.Errorf("Unexpected value: %#v (expected %#v)", str, data.encoded)
				return
			}

			err = cd.Close()
			if err != nil {
				t.Errorf("Error on close: %s", err)
				return
			}
		})
	}
}

func TestError(t *testing.T) {
	_, err := Open("INVALID_ENCODING", "INVALID_ENCODING")
	if err != EINVAL {
		t.Errorf("Unexpected error: %s (expected %s)", err, EINVAL)
	}

	cd, _ := Open("ISO-8859-15", "UTF-8")
	_, err = cd.Conv("\xc3a")
	if err != EILSEQ {
		t.Errorf("Unexpected error: %s (expected %s)", err, EILSEQ)
	}
}

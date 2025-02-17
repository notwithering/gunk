package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/alecthomas/kingpin/v2"
)

func main() {
	encodingsList := listEncodings()

	var decoding string
	kingpin.Flag("decode", "Decode into utf-8 from given encoding").PlaceHolder("ENCODING").Short('d').EnumVar(&decoding, encodingsList...)

	var listEncodings bool
	kingpin.Flag("list-encodings", "List all availible encodings").BoolVar(&listEncodings)

	var encoding string
	kingpin.Arg("encoding", "Encoding to encode stdin into").EnumVar(&encoding, encodingsList...)

	kingpin.Parse()

	if listEncodings {
		var fullNames []string
		for full := range encodings {
			fullNames = append(fullNames, full)
		}
		slices.Sort(fullNames)

		for _, full := range fullNames {
			others := encodings[full]
			fmt.Printf("%s: %s\n", full, strings.Join(others, ", "))
		}
		return
	}

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		kingpin.Fatalf("error while reading stdin:", err.Error())
	}

	if decoding != "" {
		b, err = decode(decoding, string(b))
		if err != nil {
			kingpin.Fatalf("error while decoding string: %s", err)
		}
	}

	if encoding != "" {
		s, err := encode(encoding, b)
		if err != nil {
			kingpin.Fatalf("error while encoding string: ", err)
		}
		b = []byte(s)
	}

	os.Stdout.Write(b)
}

package netcat

import (
	"io"
	"log"
)

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)

	if err != nil {
		log.Fatal(err)
	}
}

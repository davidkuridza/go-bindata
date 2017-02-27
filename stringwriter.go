// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a Modified
// BSD License that can be found in the LICENSE file.

package bindata

import (
	"fmt"
	"io"
)

type stringWriter struct {
	io.Writer
	Indent string
	WrapAt int
	c      int
}

func (w *stringWriter) Write(p []byte) (n int, err error) {
	var buf [4]byte
	buf[0] = '\\'
	buf[1] = 'x'

	for _, b := range p {
		const lowerHex = "0123456789abcdef"
		buf[2] = lowerHex[b/16]
		buf[3] = lowerHex[b%16]

		if _, err = w.Writer.Write(buf[:]); err != nil {
			return
		}

		n++
		w.c++

		if w.WrapAt == 0 || w.c%w.WrapAt != 0 {
			continue
		}

		if _, err = fmt.Fprintf(w.Writer, "\" +\n%s\"", w.Indent); err != nil {
			return
		}
	}

	return
}

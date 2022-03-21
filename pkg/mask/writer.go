package mask

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type MaskedWriter struct {
	replacements *Masks
	source       io.Reader
	target       io.Writer
}

func (w *MaskedWriter) Write() {
	scanner := bufio.NewScanner(w.source)

	for scanner.Scan() {
		fmt.Fprintln(w.target, maskLine(scanner.Text(), w.replacements))
	}
}

func NewMaskedWriter(r *Masks, s io.Reader, t io.Writer) *MaskedWriter {
	return &MaskedWriter{
		replacements: r,
		source:       s,
		target:       t,
	}
}

func maskLine(v string, r *Masks) string {
	for _, mask := range r.Values {
		ex := regexp.MustCompile(fmt.Sprintf(`(?i)%s`, mask))
		v = ex.ReplaceAllString(v, strings.Repeat(r.MaskChar, len(mask)))
	}
	return v
}

// Nicesize caps the maximum unit to EB (Exabyte, max 2^64 bytes)
package nicesize

import (
	"fmt"
	"math"
)

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

func FormatWithSpacer(size uint64, spacer string) string {
	st := "B"
	nicesize := float64(size)

	switch {
	case size < KB:
		break
	case size < MB:
		st = "KB"
		nicesize = nicesize / KB
	case size < GB:
		st = "MB"
		nicesize = nicesize / MB
	case size < TB:
		st = "GB"
		nicesize = nicesize / GB
	case size < PB:
		st = "TB"
		nicesize = nicesize / TB
	case size < EB:
		st = "PB"
		nicesize = nicesize / PB
	case size >= EB:
		st = "EB"
		nicesize = nicesize / EB

	}

	if math.Mod(nicesize, 1) < 0.01 {
		return fmt.Sprintf("%.0f%s%s", nicesize, spacer, st)
	}

	return fmt.Sprintf("%.2f%s%s", nicesize, spacer, st)
}

func Format(size uint64) string {
	return FormatWithSpacer(size, "")
}

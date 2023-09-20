package nicesize

import "fmt"

func FormatSizeWithSpacer(size int64, spacer string) string {
	const KB = 1 << 10
	const MB = KB << 10
	const GB = MB << 10
	const TB = GB << 10
	const PB = TB << 10
	const EB = PB << 10

	st := "B"
	nicesize := float64(size)

	switch {
	case size < MB:
		st = "KB"
		nicesize = float64(size) / float64(KB)
	case size < GB:
		st = "MB"
		nicesize = float64(size) / float64(MB)
	case size < TB:
		st = "GB"
		nicesize = float64(size) / float64(GB)

	case size < PB:
		st = "TB"
		nicesize = float64(size) / float64(TB)
	case size < EB:
		st = "PB"
		nicesize = float64(size) / float64(PB)
	case size >= EB:
		st = "EB"
		nicesize = float64(size) / float64(EB)

	}
	return fmt.Sprintf("%.2f%s%s", nicesize, spacer, st)
}

func FormatSize(size int64) string {
	return FormatSizeWithSpacer(size, "")
}

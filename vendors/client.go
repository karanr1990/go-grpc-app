package vendors

import "os"

type Client interface {
	TranslateText(text []string, sl string, tl string) ([]string, error)

	TranslateFile(file *os.File, sl string, tl string) (*os.File, error)
}

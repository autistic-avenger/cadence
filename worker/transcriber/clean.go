package transcriber

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CleanSubtitiles(outputDir string) error {
	subtitleFP := filepath.Join(outputDir, "video.vtt")

	content, err := os.ReadFile(subtitleFP)
	if err != nil {
		return err
	}

	cleanedStr := string(content)

	cleanedStr = strings.ReplaceAll(cleanedStr, "<c>", "")
	cleanedStr = strings.ReplaceAll(cleanedStr, "</c>", "")

	timeStampRegex := regexp.MustCompile(`<\d{2}:\d{2}:\d{2}\.\d{3}>`)
	cleanedStr = timeStampRegex.ReplaceAllString(cleanedStr, "")

	alignRegex := regexp.MustCompile(`align:start position:\d{1,2}%`)
	cleanedStr = alignRegex.ReplaceAllString(cleanedStr,"")

	return os.WriteFile(subtitleFP, []byte(cleanedStr), 0644)

}
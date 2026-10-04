package downloader

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lrstanley/go-ytdlp"
)

func DownloadSubtitle(outputDir string, url string) error {
	ytdlp.MustInstall(context.TODO(), nil)
	
	dl := ytdlp.New().
		SkipDownload().
		WriteSubs().
		WriteAutoSubs().
		ConvertSubs("vtt").
		CookiesFromBrowser("firefox").
		Output(filepath.Join(outputDir, "video.vtt"))

	_,err := dl.Run(context.TODO(),url)
	if err!=nil{
		return err
	}

	filenames ,err := filepath.Glob(filepath.Join(outputDir,"video.vtt.*.vtt"))
	if err!=nil{
		return err
	}

	if len(filenames) == 0{
		return fmt.Errorf("Subtitles are missing!")
	}

	err = os.Rename(
		filenames[0],
		filepath.Join(outputDir, "video.vtt"),
	)
	if err!=nil{
		return err
	}

	return nil
}
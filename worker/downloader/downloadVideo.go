package downloader

import (
	"context"
	"path/filepath"

	"github.com/lrstanley/go-ytdlp"
)

func DownloadVideo(url string,jobId string) error {
	ytdlp.MustInstall(context.TODO(),nil)


	outputDir,err := filepath.Abs("./jobDataStore/"+jobId)
	if err!=nil{
		return err
	}

	dl := ytdlp.New().
		FormatSort("res,ext:mp4:m4a").
		RecodeVideo("mp4").
		Output(outputDir + "/video.%(ext)s")
	
	_,err = dl.Run(context.TODO(),url)

	if err!=nil{
		return err
	}
	return nil	
}
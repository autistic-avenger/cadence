package transcriber

import (
	"cadence-worker/downloader"
	"path/filepath"
)

func TransCribe(jobID string,url string) error {
	outputDir,err := filepath.Abs("./jobDataStore/"+jobID)
	if err!=nil{
		return err
	}

	err = downloader.DownloadSubtitle(outputDir,url)
	if err!=nil{
		return err
	}
	
	return nil
}
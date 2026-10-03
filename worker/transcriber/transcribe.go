package transcriber

import (
	"os/exec"
	"path/filepath"
)

func TransCribe(jobID string) error {
	IODir,err := filepath.Abs("./jobDataStore/"+jobID)
	if err!=nil{
		return err
	}

	IOpath := IODir+"/video.mp4"
	cmd := exec.Command("whisper",IOpath,"--model","tiny","--output_dir",IODir,"--output_format","srt")
	err = cmd.Run()
	if err!=nil{
		return err
	}
	return nil
}
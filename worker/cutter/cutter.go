package cutter

import (
	"cadence-worker/picker"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func MakeClips(jobID string, clips picker.ClipResponse) error {

	err := os.MkdirAll(filepath.Join("jobDataStore",jobID,"clips"),0644)
	if err!=nil{
		return err
	}

	inputPath := "./jobDataStore/" + jobID + "/video.mp4"
	for i,clip := range clips.Clips{
		outputPath := fmt.Sprintf("./jobDataStore/%s/clips/clip%d.mp4",jobID,i+1)
		
		cmd := exec.Command("ffmpeg",
			"-y",
			"-ss", clip.Start,
			"-to", clip.End,
			"-i", inputPath,
			"-c", "copy",
			"-avoid_negative_ts", "make_zero",
			outputPath,
		)

		err := cmd.Run()
		if err!=nil{
			return err
		}
		clip.Location = fmt.Sprintf("clip%d.mp4",i+1)
		fmt.Println("[CUTTING] :",i+1,"/",len(clips.Clips))
	}
	return nil
}
package picker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
)

type Clip struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Title string `json:"title"`
}

type ClipResponse struct {
	Clips []Clip `json:"clips"`
}

func extractJSON(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "```json") {
		s = strings.TrimPrefix(s, "```json")
	} else if strings.HasPrefix(s, "```") {
		s = strings.TrimPrefix(s, "```")
	}
	s = strings.TrimSuffix(s, "```")
	s = strings.TrimSpace(s)

	return s
}

func PickClips(jobID string) (ClipResponse, error) {
	vttPath := "./jobDataStore/" + jobID + "/video.vtt"
	transcript, err := os.ReadFile(vttPath)
	if err != nil {
		return ClipResponse{}, fmt.Errorf("failed to read transcript file %s: %w", vttPath, err)
	}

	apiKey := strings.TrimSpace(os.Getenv("OPENROUTER_API_KEY"))
	if apiKey == "" {
		return ClipResponse{}, fmt.Errorf("OPENROUTER_API_KEY is not in .env")
	}

	openS := openrouter.New(
		openrouter.WithSecurity(apiKey),
	)

	llmModel := strings.TrimSpace(os.Getenv("OPENROUTER_MODEL"))
	if llmModel == "" {
		llmModel = "google/gemma-4-26b-a4b-it:free"
	}

	batchSize := 300

	maxToken := int64(3000)

	getSysPrompt, err := os.ReadFile("./picker/systemPrompt.txt")
	if err != nil {
		return ClipResponse{}, fmt.Errorf("failed to read system prompt: %w", err)
	}
	systemPrompt := string(getSysPrompt)

	lines := strings.Split(string(transcript), "\n")

	clipsArray := ClipResponse{
		Clips: make([]Clip, 0),
	}

	var nonEmptyLines []string
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed != "" {
			nonEmptyLines = append(nonEmptyLines, trimmed)
		}
	}

	if len(nonEmptyLines) == 0 {
		return ClipResponse{}, fmt.Errorf("transcript is empty!")
	}

	totalBatches := (len(nonEmptyLines) + batchSize - 1) / batchSize

	successfulBatches := 0
	for i := 0; i < len(nonEmptyLines); i += batchSize {
		end := i + batchSize
		if end > len(nonEmptyLines) {
			end = len(nonEmptyLines)
		}

		batchLines := nonEmptyLines[i:end]
		batchText := strings.Join(batchLines, "\n")
		if strings.TrimSpace(batchText) == "" {
			continue
		}

		batchNum := (i / batchSize) + 1
		fmt.Printf("[PICKING] Processing batch %d/%d\n", batchNum, totalBatches)

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		res, err := openS.Chat.Send(ctx, components.ChatRequest{
			Model:     &llmModel,
			MaxTokens: optionalnullable.From(&maxToken),
			Messages: []components.ChatMessages{
				components.CreateChatMessagesSystem(
					components.ChatSystemMessage{
						Content: components.CreateChatSystemMessageContentStr(
							systemPrompt,
						),
						Role: components.ChatSystemMessageRoleSystem,
					},
				),
				components.CreateChatMessagesUser(
					components.ChatUserMessage{
						Content: components.CreateChatUserMessageContentStr(
							"Here is the transcript:\n" + batchText,
						),
						Role: components.ChatUserMessageRoleUser,
					},
				),
			},
		}, nil)
		cancel()

		if err != nil {
			fmt.Printf("[BATCH FAIL] Batch %d/%d\n", batchNum, totalBatches)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if res.ChatResult == nil || len(res.ChatResult.Choices) == 0 {
			fmt.Printf("[MODEL ERROR] Batch %d/%d\n", batchNum, totalBatches)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		llmContent, ok := res.ChatResult.Choices[0].Message.Content.Get()
		if !ok || llmContent == nil || llmContent.Str == nil {
			fmt.Printf("[LLM CONTENT ERROR] Batch %d/%d\n", batchNum, totalBatches)
			time.Sleep(1 * time.Second)
			continue
		}

		clipsJson := extractJSON(*llmContent.Str)

		var batchResp ClipResponse
		if err := json.Unmarshal([]byte(clipsJson), &batchResp); err != nil {
			var rawClips []Clip
			if err2 := json.Unmarshal([]byte(clipsJson), &rawClips); err2 == nil {
				batchResp.Clips = rawClips
			} else {
				fmt.Printf("[PICKING WARNING] Batch %d/%d json unmarshal failed: %v. Raw text: %s\n", batchNum, totalBatches, err, clipsJson)
				time.Sleep(1 * time.Second)
				continue
			}
		}

		successfulBatches++
		if len(batchResp.Clips) > 0 {
			clipsArray.Clips = append(clipsArray.Clips, batchResp.Clips...)
		}

		time.Sleep(1 * time.Second)
	}

	if successfulBatches == 0 {
		return ClipResponse{}, fmt.Errorf("failed to process any batch")
	}

	fmt.Printf("[PICKING] CLIPS Collected:%d\n", len(clipsArray.Clips))
	return clipsArray, nil
}



func CreateJson(jobID string,clips ClipResponse)error {
 	clipsPath := "./jobDataStore/" + jobID + "/clips.json"

	jsonByteStream,err := json.Marshal(clips.Clips)
	if err!=nil{
		return err
	}
	err = os.WriteFile(clipsPath,jsonByteStream,0644)
	if err!=nil{
		return err
	}

	return nil 
}
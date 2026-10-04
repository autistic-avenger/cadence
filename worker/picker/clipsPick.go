package picker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

type Clip struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Title string `json:"title"`
}

type ClipResponse struct {
	Clips []Clip `json:"clips"`
}


func PickClips(jobID string) (ClipResponse, error) {
	transcript, err := os.ReadFile("./jobDataStore/" + jobID + "/video.vtt")
	if err != nil {
		return ClipResponse{}, err
	}

	openS := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	llmModel := "google/gemini-2.5-flash-lite"


	if err!=nil{
		return ClipResponse{},err
	}
	getSysPrompt,err := os.ReadFile("./picker/systemPrompt.txt")
	if err!=nil{
		return ClipResponse{},err
	}

	systemPrompt := string(getSysPrompt)

	ctx := context.Background()
	res, err := openS.Chat.Send(ctx, components.ChatRequest{
		Model: &llmModel,
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
						"Here is the transcript:\n"+ string(transcript),
					),
					Role: components.ChatUserMessageRoleUser,
				},
			),
		},
	}, nil)

	if err!=nil{
		return ClipResponse{},err
	}

	llmContent,ok := res.ChatResult.Choices[0].Message.Content.Get()

	if !ok{
		return ClipResponse{},fmt.Errorf("Model did not return Content!")
	}
	var clipsArray ClipResponse
	
	clipsJson := *llmContent.Str
	if clipsJson != ""{
		clipsJson = strings.TrimPrefix(clipsJson,"```json")
		clipsJson = strings.TrimSuffix(clipsJson,"```")
	}

	err = json.Unmarshal([]byte(clipsJson),&clipsArray)
	if err!=nil{
		return ClipResponse{},err
	}

	fmt.Println(clipsJson) //retard Don't forget to remove this after testing
	

	return clipsArray,nil
}
package usecase

import (
	"context"
	"fmt"

	"ai-task/dto"

	"github.com/sashabaranov/go-openai"
)

type LaptopRecomendedUsecase interface {
	LaptopRecomended(request dto.RequestData, key string) (string, error)
}

type Usecase struct{}

func NewUseCase() LaptopRecomendedUsecase {
	return &Usecase{}
}

func (uc *Usecase) getCompletionMessages(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func (us *Usecase) LaptopRecomended(request dto.RequestData, key string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(key)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "You must give a laptop recommendation for user with specific in formation about the right price and specification from user like Brand, RAM, Display, Processor, Display and Storage type",
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("Laptop recommendation with budget : RP%s for %s. with brand : %s", request.Budget,request.Purpose,request.Brand),
		},
	}

	response, err := us.getCompletionMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := response.Choices[0].Message.Content
	return answer, nil
}

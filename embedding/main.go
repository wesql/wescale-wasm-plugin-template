package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strings"
)

func main() {
	TOKEN := os.Getenv("OPENAI_API_KEY")

	config := openai.DefaultConfig(TOKEN)
	config.BaseURL = "https://api.gptsapi.net/v1"

	client := openai.NewClientWithConfig(config)

	// Create an EmbeddingRequest for the user query
	queryReq := openai.EmbeddingRequest{
		Input: []string{"How many chucks would a woodchuck chuck"},
		Model: openai.AdaEmbeddingV2,
	}

	// Create an embedding for the user query
	queryResponse, err := client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		log.Fatal("Error creating query embedding:", err)
	}

	// Create an EmbeddingRequest for the target text
	targetReq := openai.EmbeddingRequest{
		Input: []string{"How many chucks would a woodchuck chuck if the woodchuck could chuck wood"},
		Model: openai.AdaEmbeddingV2,
	}

	// Create an embedding for the target text
	targetResponse, err := client.CreateEmbeddings(context.Background(), targetReq)
	if err != nil {
		log.Fatal("Error creating target embedding:", err)
	}

	// Now that we have the embeddings for the user query and the target text, we
	// can calculate their similarity.
	queryEmbedding := queryResponse.Data[0]
	targetEmbedding := targetResponse.Data[0]

	similarity, err := queryEmbedding.DotProduct(&targetEmbedding)
	if err != nil {
		log.Fatal("Error calculating dot product:", err)
	}

	log.Printf("The similarity score between the query and the target is %f", similarity)

	fmt.Println(embedding("How many chucks would a woodchuck chuck"))
}

func embedding(input ...string) string {
	TOKEN := os.Getenv("OPENAI_API_KEY")

	config := openai.DefaultConfig(TOKEN)
	config.BaseURL = "https://api.gptsapi.net/v1"

	client := openai.NewClientWithConfig(config)

	// Create an EmbeddingRequest for the user query
	queryReq := openai.EmbeddingRequest{
		Input: input,
		Model: openai.AdaEmbeddingV2,
	}

	// Create an embedding for the user query
	queryResponse, err := client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		log.Fatal("Error creating query embedding:", err)
	}

	queryEmbedding := queryResponse.Data[0]

	// convert []float32 to string
	return Float32SliceToString(queryEmbedding.Embedding)
}

func Float32SliceToString(slice []float32) string {
	strs := make([]string, len(slice))
	for i, v := range slice {
		strs[i] = fmt.Sprintf("%f", v)
	}
	return strings.Join(strs, ",")
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sedletsky-f5/langchaingo/agents"
	"github.com/sedletsky-f5/langchaingo/chains"
	"github.com/sedletsky-f5/langchaingo/llms/openai"
	"github.com/sedletsky-f5/langchaingo/tools"
	"github.com/sedletsky-f5/langchaingo/tools/serpapi"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	llm, err := openai.New()
	if err != nil {
		return err
	}
	search, err := serpapi.New()
	if err != nil {
		return err
	}
	agentTools := []tools.Tool{
		tools.Calculator{},
		search,
	}
	executor, err := agents.Initialize(
		llm,
		agentTools,
		agents.ZeroShotReactDescription,
		agents.WithMaxIterations(3),
	)
	if err != nil {
		return err
	}
	question := "Who is Olivia Wilde's boyfriend? What is his current age raised to the 0.23 power?"
	answer, err := chains.Run(context.Background(), executor, question)
	fmt.Println(answer)
	return err
}

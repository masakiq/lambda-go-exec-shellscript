package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"os/exec"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
}

func hello(event MyEvent) (string, error) {
	out, err := exec.Command("./hello-bash", event.Name).Output()
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return fmt.Sprintf(string(out[:])), nil
}

func main() {
	lambda.Start(hello)
}

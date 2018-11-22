package main

import (
	"context"
	"fmt"
	"github.com/sample-web/client/common"
)

func main() {
	//response, err := runGetFullName()
	//response, err := runPostMessage()
	response, err := runSetCookie()
	//response, err := runBasicAuth()
	if err != nil {
		fmt.Println("response error,,,,,,")
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%+v\n", response)
}

func runGetFullName() (string, error) {
	cli, err := common.NewClient("http://localhost:8080")
	if err != nil {
		return "NewClient error", err
	}

	return cli.GetFullName(context.Background(), "hogehoge", "hugahuga")
}

func runPostMessage() (string, error) {
	cli, err := common.NewClient("http://localhost:8080")
	if err != nil {
		return "NewClient error", err
	}

	return cli.PostMessage(context.Background(), "Chester", "sample message")
}

func runSetCookie() (string, error) {
	cli, err := common.NewClient("http://localhost:8080")
	if err != nil {
		return "NewClient error", err
	}

	return cli.SendRequest(context.Background(), "GET", "/set-cookie")
}

func runBasicAuth() (string, error) {
	cli, err := common.NewClient("http://localhost:8080", common.BasicAuth("admin", "secret"))
	if err != nil {
		return "NewClient error", err
	}

	return cli.SendRequest(context.Background(), "GET", "/auth/admin")
}
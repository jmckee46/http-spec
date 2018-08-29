package main

import (
	"fmt"
	"os"
	"strings"
)

func requestFromFile(context *context) (*request, error) {
	message, err := messageFromFile(context)

	if err != nil {
		return nil, err
	}

	return &request{
		message,
		context.Hostname,
		context.Scheme,
	}, err
}

type request struct {
	*message
	Hostname string
	Scheme   string
}

func (request *request) Method() string {
	return strings.Split(request.FirstLine.Text, " ")[0]
}

func (request *request) uri() string {
	return strings.Split(request.FirstLine.Text, " ")[1]
}

func (request *request) URL() string {
	if request.Scheme == "" || request.Hostname == "" {
		fmt.Fprintf(
			os.Stderr,
			"%s%s%s%s\n",
			Yellow,
			"DEPRECATION: absolute (proxy) URLs support will be removed in 1.0\n",
			"DEPRECATION: use -hostname and -scheme cli options instead",
			Reset,
		)

		return request.uri()
	}

	return request.Scheme + "://" + request.Hostname + request.uri()
}

func (request *request) Version() string {
	return strings.Split(request.FirstLine.Text, " ")[2]
}

func (request *request) String() string {
	lineStrings := []string{}

	lineStrings = append(lineStrings, request.FirstLine.String())

	for _, l := range request.HeaderLines {
		lineStrings = append(lineStrings, l.String())
	}

	lineStrings = append(lineStrings, request.BlankLine.String())

	if request.BodyLines != nil {
		for _, l := range request.BodyLines {
			lineStrings = append(lineStrings, l.String())
		}
	}

	return strings.Join(lineStrings, "\n")
}

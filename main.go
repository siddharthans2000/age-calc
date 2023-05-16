package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Date struct {
	Date  int `json:"date"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var date Date
	var msg string
	var no_days int
	var no_months int
	var no_years int

	t := time.Now()
	days_in_month := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	err := json.Unmarshal([]byte(event.Body), &date)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if t.Year() < date.Year {
		msg = fmt.Sprintf("Please enter a valid age !")
	} else {

		no_days = days_in_month[t.Month()] - t.Day() + date.Date
		no_months = (12 - date.Month) + int(t.Month()) - 1 + (no_days)/30
		no_years = (t.Year() - date.Year - 1) + (no_months)/12
		no_months = no_months % 12
		no_days = no_days % 30

		msg = fmt.Sprintf("You are %d years old %d months old %d days old", no_years, no_months, no_days)
	}
	responseBody := Response{
		Message: msg,
	}
	jbytes, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jbytes),
	}

	return response, err
}

func main() {
	lambda.Start(HandleRequest)
}

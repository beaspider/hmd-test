package main

import (
	"strings"
	"testing"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/dynamodb"
)

func TestWorkInProgress(t *testing.T) {

	// This assumes you have ENV vars: AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
	auth, err := aws.EnvAuth()
	if err != nil {
		t.Error("Error :", err)
	}

	ddbs := dynamodb.Server{auth, aws.EUWest}
	response, err := ddbs.ListTables()
	if err != nil {
		t.Error("Error :", err)
	}

	t.Logf("The tablenames are %s!", strings.Join(response, ", "))

	t.Log("TODO: look at DynamoDB Query and Scan")
	// http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html

}

package usage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"testing"
)

func TestUsageQueryCommand(t *testing.T) {
	command := usageQueryCommand

	path, err := utils.GetResourceDefinition(command.Path)
	if err != nil {
		t.Fatalf("Received error getting path structure from swagger, err: %v", err)
	}

	method := path.GetMethod(command.Method)
	if method == nil {
		t.Fatalf("Failed to retrieve method definition from swagger")
	}

	if command.Description != "" {
		if command.Description != method.Description {
			t.Errorf("Description doesn't match swagger. got: %v, want :%v", command.Description, method.Description)
		}
	}
}

func TestUsageQueryResultsCommand(t *testing.T) {
	command := usageQueryResultsCommand

	path, err := utils.GetResourceDefinition(command.Path)
	if err != nil {
		t.Fatalf("Received error getting path structure from swagger, err: %v", err)
	}

	method := path.GetMethod(command.Method)
	if method == nil {
		t.Fatalf("Failed to retrieve method definition from swagger")
	}

	if command.Description != "" {
		if command.Description != method.Description {
			t.Errorf("Description doesn't match swagger. got: %v, want :%v", command.Description, method.Description)
		}
	}
}
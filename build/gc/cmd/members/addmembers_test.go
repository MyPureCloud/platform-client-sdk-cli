package members

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"testing"
)

func TestAddMembersCommand(t *testing.T) {
	command := addMembersCommand

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

func TestGetMembersCommand(t *testing.T) {
	command := getMembersCommand

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

package members

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"testing"
)

func TestAddMembersOperation(t *testing.T) {
	utils.TestAgainstSwaggerDefinition(t, addMembersOperation)
}

func TestGetMembersOperation(t *testing.T) {
	utils.TestAgainstSwaggerDefinition(t, getMembersOperation)
}

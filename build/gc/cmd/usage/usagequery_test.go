package usage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"testing"
)

func TestUsageQueryOperation(t *testing.T) {
	utils.TestAgainstSwaggerDefinition(t, usageQueryOperation)
}

func TestUsageQueryResultsOperation(t *testing.T) {
	utils.TestAgainstSwaggerDefinition(t, usageQueryResultsOperation)
}

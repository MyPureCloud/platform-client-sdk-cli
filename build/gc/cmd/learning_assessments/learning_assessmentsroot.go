package learning_assessments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assessments_scoring"
)

func init() {
	learning_assessmentsCmd.AddCommand(learning_assessments_scoring.Cmdlearning_assessments_scoring())
	learning_assessmentsCmd.Short = utils.GenerateCustomDescription(learning_assessmentsCmd.Short, learning_assessments_scoring.Description, )
	learning_assessmentsCmd.Long = learning_assessmentsCmd.Short
}

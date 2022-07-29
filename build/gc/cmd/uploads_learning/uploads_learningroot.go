package uploads_learning

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/uploads_learning_coverart"
)

func init() {
	uploads_learningCmd.AddCommand(uploads_learning_coverart.Cmduploads_learning_coverart())
	uploads_learningCmd.Short = utils.GenerateCustomDescription(uploads_learningCmd.Short, uploads_learning_coverart.Description, )
	uploads_learningCmd.Long = uploads_learningCmd.Short
}

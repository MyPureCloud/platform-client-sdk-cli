package uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/uploads_workforcemanagement"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/uploads_publicassets"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/uploads_recordings"
)

func init() {
	uploadsCmd.AddCommand(uploads_workforcemanagement.Cmduploads_workforcemanagement())
	uploadsCmd.AddCommand(uploads_publicassets.Cmduploads_publicassets())
	uploadsCmd.AddCommand(uploads_recordings.Cmduploads_recordings())
	uploadsCmd.Short = utils.GenerateCustomDescription(uploadsCmd.Short, uploads_workforcemanagement.Description, uploads_publicassets.Description, uploads_recordings.Description, )
	uploadsCmd.Long = uploadsCmd.Short
}

package recording_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_uploads_reports"
)

func init() {
	recording_uploadsCmd.AddCommand(recording_uploads_reports.Cmdrecording_uploads_reports())
	recording_uploadsCmd.Short = utils.GenerateCustomDescription(recording_uploadsCmd.Short, recording_uploads_reports.Description, )
	recording_uploadsCmd.Long = recording_uploadsCmd.Short
}

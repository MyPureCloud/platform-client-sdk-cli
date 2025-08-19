package recordings_retention

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_retention_query"
)

func init() {
	recordings_retentionCmd.AddCommand(recordings_retention_query.Cmdrecordings_retention_query())
	recordings_retentionCmd.Short = utils.GenerateCustomDescription(recordings_retentionCmd.Short, recordings_retention_query.Description, )
	recordings_retentionCmd.Long = recordings_retentionCmd.Short
}

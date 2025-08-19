package scripts_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts_uploads_status"
)

func init() {
	scripts_uploadsCmd.AddCommand(scripts_uploads_status.Cmdscripts_uploads_status())
	scripts_uploadsCmd.Short = utils.GenerateCustomDescription(scripts_uploadsCmd.Short, scripts_uploads_status.Description, )
	scripts_uploadsCmd.Long = scripts_uploadsCmd.Short
}

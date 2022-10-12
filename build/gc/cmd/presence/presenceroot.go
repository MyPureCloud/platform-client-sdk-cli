package presence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_sources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_users"
)

func init() {
	presenceCmd.AddCommand(presence_sources.Cmdpresence_sources())
	presenceCmd.AddCommand(presence_users.Cmdpresence_users())
	presenceCmd.Short = utils.GenerateCustomDescription(presenceCmd.Short, presence_sources.Description, presence_users.Description, )
	presenceCmd.Long = presenceCmd.Short
}

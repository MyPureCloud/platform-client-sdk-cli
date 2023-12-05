package presence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_definitions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_sources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_users"
)

func init() {
	presenceCmd.AddCommand(presence_definitions.Cmdpresence_definitions())
	presenceCmd.AddCommand(presence_settings.Cmdpresence_settings())
	presenceCmd.AddCommand(presence_sources.Cmdpresence_sources())
	presenceCmd.AddCommand(presence_users.Cmdpresence_users())
	presenceCmd.Short = utils.GenerateCustomDescription(presenceCmd.Short, presence_definitions.Description, presence_settings.Description, presence_sources.Description, presence_users.Description, )
	presenceCmd.Long = presenceCmd.Short
}

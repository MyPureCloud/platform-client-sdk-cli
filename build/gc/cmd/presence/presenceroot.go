package presence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_sources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_definitions"
)

func init() {
	presenceCmd.AddCommand(presence_settings.Cmdpresence_settings())
	presenceCmd.AddCommand(presence_sources.Cmdpresence_sources())
	presenceCmd.AddCommand(presence_users.Cmdpresence_users())
	presenceCmd.AddCommand(presence_definitions.Cmdpresence_definitions())
	presenceCmd.Short = utils.GenerateCustomDescription(presenceCmd.Short, presence_settings.Description, presence_sources.Description, presence_users.Description, presence_definitions.Description, )
	presenceCmd.Long = presenceCmd.Short
}

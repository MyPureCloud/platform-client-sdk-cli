package webchat

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_guest"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_deployments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_settings"
)

func init() {
	webchatCmd.AddCommand(webchat_guest.Cmdwebchat_guest())
	webchatCmd.AddCommand(webchat_deployments.Cmdwebchat_deployments())
	webchatCmd.AddCommand(webchat_settings.Cmdwebchat_settings())
	webchatCmd.Short = utils.GenerateCustomDescription(webchatCmd.Short, webchat_guest.Description, webchat_deployments.Description, webchat_settings.Description, )
	webchatCmd.Long = webchatCmd.Short
}

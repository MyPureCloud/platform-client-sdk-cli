package webchat

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_guest"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_deployments"
)

func init() {
	webchatCmd.AddCommand(webchat_guest.Cmdwebchat_guest())
	webchatCmd.AddCommand(webchat_settings.Cmdwebchat_settings())
	webchatCmd.AddCommand(webchat_deployments.Cmdwebchat_deployments())
	webchatCmd.Short = utils.GenerateCustomDescription(webchatCmd.Short, webchat_guest.Description, webchat_settings.Description, webchat_deployments.Description, )
	webchatCmd.Long = webchatCmd.Short
}

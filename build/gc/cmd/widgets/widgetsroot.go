package widgets

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/widgets_deployments"
)

func init() {
	widgetsCmd.AddCommand(widgets_deployments.Cmdwidgets_deployments())
	widgetsCmd.Short = utils.GenerateCustomDescription(widgetsCmd.Short, widgets_deployments.Description, )
	widgetsCmd.Long = widgetsCmd.Short
}

package architect_systemprompts_resources

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_systemprompts_resources_uploads"
)

func init() {
	architect_systemprompts_resourcesCmd.AddCommand(architect_systemprompts_resources_uploads.Cmdarchitect_systemprompts_resources_uploads())
	architect_systemprompts_resourcesCmd.Short = utils.GenerateCustomDescription(architect_systemprompts_resourcesCmd.Short, architect_systemprompts_resources_uploads.Description, )
	architect_systemprompts_resourcesCmd.Long = architect_systemprompts_resourcesCmd.Short
}

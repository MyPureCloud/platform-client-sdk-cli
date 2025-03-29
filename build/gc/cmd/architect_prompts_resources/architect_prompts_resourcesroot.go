package architect_prompts_resources

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_prompts_resources_audio"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_prompts_resources_uploads"
)

func init() {
	architect_prompts_resourcesCmd.AddCommand(architect_prompts_resources_audio.Cmdarchitect_prompts_resources_audio())
	architect_prompts_resourcesCmd.AddCommand(architect_prompts_resources_uploads.Cmdarchitect_prompts_resources_uploads())
	architect_prompts_resourcesCmd.Short = utils.GenerateCustomDescription(architect_prompts_resourcesCmd.Short, architect_prompts_resources_audio.Description, architect_prompts_resources_uploads.Description, )
	architect_prompts_resourcesCmd.Long = architect_prompts_resourcesCmd.Short
}

package scripts_published

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts_published_pages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts_published_variables"
)

func init() {
	scripts_publishedCmd.AddCommand(scripts_published_pages.Cmdscripts_published_pages())
	scripts_publishedCmd.AddCommand(scripts_published_variables.Cmdscripts_published_variables())
	scripts_publishedCmd.Short = utils.GenerateCustomDescription(scripts_publishedCmd.Short, scripts_published_pages.Description, scripts_published_variables.Description, )
	scripts_publishedCmd.Long = scripts_publishedCmd.Short
}

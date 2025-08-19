package scripts_published_divisionviews

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts_published_divisionviews_variables"
)

func init() {
	scripts_published_divisionviewsCmd.AddCommand(scripts_published_divisionviews_variables.Cmdscripts_published_divisionviews_variables())
	scripts_published_divisionviewsCmd.Short = utils.GenerateCustomDescription(scripts_published_divisionviewsCmd.Short, scripts_published_divisionviews_variables.Description, )
	scripts_published_divisionviewsCmd.Long = scripts_published_divisionviewsCmd.Short
}

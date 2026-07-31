package responsemanagement_responses_divisionviews

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_responses_divisionviews_query"
)

func init() {
	responsemanagement_responses_divisionviewsCmd.AddCommand(responsemanagement_responses_divisionviews_query.Cmdresponsemanagement_responses_divisionviews_query())
	responsemanagement_responses_divisionviewsCmd.Short = utils.GenerateCustomDescription(responsemanagement_responses_divisionviewsCmd.Short, responsemanagement_responses_divisionviews_query.Description, )
	responsemanagement_responses_divisionviewsCmd.Long = responsemanagement_responses_divisionviewsCmd.Short
}

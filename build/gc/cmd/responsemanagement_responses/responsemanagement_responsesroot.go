package responsemanagement_responses

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_responses_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_responses_divisionviews"
)

func init() {
	responsemanagement_responsesCmd.AddCommand(responsemanagement_responses_query.Cmdresponsemanagement_responses_query())
	responsemanagement_responsesCmd.AddCommand(responsemanagement_responses_divisionviews.Cmdresponsemanagement_responses_divisionviews())
	responsemanagement_responsesCmd.Short = utils.GenerateCustomDescription(responsemanagement_responsesCmd.Short, responsemanagement_responses_query.Description, responsemanagement_responses_divisionviews.Description, )
	responsemanagement_responsesCmd.Long = responsemanagement_responsesCmd.Short
}

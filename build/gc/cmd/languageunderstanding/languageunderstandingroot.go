package languageunderstanding

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains"
)

func init() {
	languageunderstandingCmd.AddCommand(languageunderstanding_domains.Cmdlanguageunderstanding_domains())
	languageunderstandingCmd.Short = utils.GenerateCustomDescription(languageunderstandingCmd.Short, languageunderstanding_domains.Description, )
	languageunderstandingCmd.Long = languageunderstandingCmd.Short
}

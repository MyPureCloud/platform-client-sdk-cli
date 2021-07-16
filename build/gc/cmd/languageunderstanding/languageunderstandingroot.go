package languageunderstanding

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_miners"
)

func init() {
	languageunderstandingCmd.AddCommand(languageunderstanding_domains.Cmdlanguageunderstanding_domains())
	languageunderstandingCmd.AddCommand(languageunderstanding_miners.Cmdlanguageunderstanding_miners())
	languageunderstandingCmd.Short = utils.GenerateCustomDescription(languageunderstandingCmd.Short, languageunderstanding_domains.Description, languageunderstanding_miners.Description, )
	languageunderstandingCmd.Long = languageunderstandingCmd.Short
}

package languageunderstanding

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_miners"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_settings"
)

func init() {
	languageunderstandingCmd.AddCommand(languageunderstanding_domains.Cmdlanguageunderstanding_domains())
	languageunderstandingCmd.AddCommand(languageunderstanding_miners.Cmdlanguageunderstanding_miners())
	languageunderstandingCmd.AddCommand(languageunderstanding_settings.Cmdlanguageunderstanding_settings())
	languageunderstandingCmd.Short = utils.GenerateCustomDescription(languageunderstandingCmd.Short, languageunderstanding_domains.Description, languageunderstanding_miners.Description, languageunderstanding_settings.Description, )
	languageunderstandingCmd.Long = languageunderstandingCmd.Short
}

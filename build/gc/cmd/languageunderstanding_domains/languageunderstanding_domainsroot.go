package languageunderstanding_domains

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains_feedback"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains_versions"
)

func init() {
	languageunderstanding_domainsCmd.AddCommand(languageunderstanding_domains_feedback.Cmdlanguageunderstanding_domains_feedback())
	languageunderstanding_domainsCmd.AddCommand(languageunderstanding_domains_versions.Cmdlanguageunderstanding_domains_versions())
	languageunderstanding_domainsCmd.Short = utils.GenerateCustomDescription(languageunderstanding_domainsCmd.Short, languageunderstanding_domains_feedback.Description, languageunderstanding_domains_versions.Description, )
	languageunderstanding_domainsCmd.Long = languageunderstanding_domainsCmd.Short
}

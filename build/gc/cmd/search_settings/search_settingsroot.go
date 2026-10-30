package search_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/search_settings_contentsearch"
)

func init() {
	search_settingsCmd.AddCommand(search_settings_contentsearch.Cmdsearch_settings_contentsearch())
	search_settingsCmd.Short = utils.GenerateCustomDescription(search_settingsCmd.Short, search_settings_contentsearch.Description, )
	search_settingsCmd.Long = search_settingsCmd.Short
}

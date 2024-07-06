package languages_translations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languages_translations_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languages_translations_builtin"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languages_translations_organization"
)

func init() {
	languages_translationsCmd.AddCommand(languages_translations_users.Cmdlanguages_translations_users())
	languages_translationsCmd.AddCommand(languages_translations_builtin.Cmdlanguages_translations_builtin())
	languages_translationsCmd.AddCommand(languages_translations_organization.Cmdlanguages_translations_organization())
	languages_translationsCmd.Short = utils.GenerateCustomDescription(languages_translationsCmd.Short, languages_translations_users.Description, languages_translations_builtin.Description, languages_translations_organization.Description, )
	languages_translationsCmd.Long = languages_translationsCmd.Short
}

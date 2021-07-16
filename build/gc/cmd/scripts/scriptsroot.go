package scripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts_export"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts_uploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts_pages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts_published"
)

func init() {
	scriptsCmd.AddCommand(scripts_export.Cmdscripts_export())
	scriptsCmd.AddCommand(scripts_uploads.Cmdscripts_uploads())
	scriptsCmd.AddCommand(scripts_pages.Cmdscripts_pages())
	scriptsCmd.AddCommand(scripts_published.Cmdscripts_published())
	scriptsCmd.Short = utils.GenerateCustomDescription(scriptsCmd.Short, scripts_export.Description, scripts_uploads.Description, scripts_pages.Description, scripts_published.Description, )
	scriptsCmd.Long = scriptsCmd.Short
}

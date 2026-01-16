package guides

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_sessions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_uploads"
)

func init() {
	guidesCmd.AddCommand(guides_jobs.Cmdguides_jobs())
	guidesCmd.AddCommand(guides_sessions.Cmdguides_sessions())
	guidesCmd.AddCommand(guides_versions.Cmdguides_versions())
	guidesCmd.AddCommand(guides_uploads.Cmdguides_uploads())
	guidesCmd.Short = utils.GenerateCustomDescription(guidesCmd.Short, guides_jobs.Description, guides_sessions.Description, guides_versions.Description, guides_uploads.Description, )
	guidesCmd.Long = guidesCmd.Short
}

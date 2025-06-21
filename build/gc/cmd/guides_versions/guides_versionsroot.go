package guides_versions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_versions_jobs"
)

func init() {
	guides_versionsCmd.AddCommand(guides_versions_jobs.Cmdguides_versions_jobs())
	guides_versionsCmd.Short = utils.GenerateCustomDescription(guides_versionsCmd.Short, guides_versions_jobs.Description, )
	guides_versionsCmd.Long = guides_versionsCmd.Short
}

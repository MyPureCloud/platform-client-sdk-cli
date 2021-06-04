package orphanrecordings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orphanrecordings_media"
)

func init() {
	orphanrecordingsCmd.AddCommand(orphanrecordings_media.Cmdorphanrecordings_media())
	orphanrecordingsCmd.Short = utils.GenerateCustomDescription(orphanrecordingsCmd.Short, orphanrecordings_media.Description, )
	orphanrecordingsCmd.Long = orphanrecordingsCmd.Short
}

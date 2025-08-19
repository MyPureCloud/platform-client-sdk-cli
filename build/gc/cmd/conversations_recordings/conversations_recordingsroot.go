package conversations_recordings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_recordings_annotations"
)

func init() {
	conversations_recordingsCmd.AddCommand(conversations_recordings_annotations.Cmdconversations_recordings_annotations())
	conversations_recordingsCmd.Short = utils.GenerateCustomDescription(conversations_recordingsCmd.Short, conversations_recordings_annotations.Description, )
	conversations_recordingsCmd.Long = conversations_recordingsCmd.Short
}

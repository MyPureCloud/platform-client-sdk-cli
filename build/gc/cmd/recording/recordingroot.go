package recording

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_crossplatform"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_mediaretentionpolicies"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_recordingkeys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_localkeys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_keyconfigurations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_batchrequests"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_uploads"
)

func init() {
	recordingCmd.AddCommand(recording_crossplatform.Cmdrecording_crossplatform())
	recordingCmd.AddCommand(recording_mediaretentionpolicies.Cmdrecording_mediaretentionpolicies())
	recordingCmd.AddCommand(recording_recordingkeys.Cmdrecording_recordingkeys())
	recordingCmd.AddCommand(recording_localkeys.Cmdrecording_localkeys())
	recordingCmd.AddCommand(recording_keyconfigurations.Cmdrecording_keyconfigurations())
	recordingCmd.AddCommand(recording_batchrequests.Cmdrecording_batchrequests())
	recordingCmd.AddCommand(recording_jobs.Cmdrecording_jobs())
	recordingCmd.AddCommand(recording_settings.Cmdrecording_settings())
	recordingCmd.AddCommand(recording_uploads.Cmdrecording_uploads())
	recordingCmd.Short = utils.GenerateCustomDescription(recordingCmd.Short, recording_crossplatform.Description, recording_mediaretentionpolicies.Description, recording_recordingkeys.Description, recording_localkeys.Description, recording_keyconfigurations.Description, recording_batchrequests.Description, recording_jobs.Description, recording_settings.Description, recording_uploads.Description, )
	recordingCmd.Long = recordingCmd.Short
}

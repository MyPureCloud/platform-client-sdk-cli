package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingsettingsDud struct { 
    


    


    


    


    

}

// Recordingsettings
type Recordingsettings struct { 
    // MaxSimultaneousStreams - Maximum number of simultaneous screen recording streams
    MaxSimultaneousStreams int `json:"maxSimultaneousStreams"`


    // MaxConfigurableScreenRecordingStreams - Upper limit that maxSimultaneousStreams can be configured
    MaxConfigurableScreenRecordingStreams int `json:"maxConfigurableScreenRecordingStreams"`


    // RegionalRecordingStorageEnabled - Store call recordings in the region where they are intended to be recorded, otherwise in the organization's home region
    RegionalRecordingStorageEnabled bool `json:"regionalRecordingStorageEnabled"`


    // RecordingPlaybackUrlTtl - The duration in minutes for which the generated URL for recording playback remains valid.The default duration is set to 60 minutes, with a minimum allowable duration of 2 minutes and a maximum of 60 minutes.
    RecordingPlaybackUrlTtl int `json:"recordingPlaybackUrlTtl"`


    // RecordingBatchDownloadUrlTtl - TThe duration in minutes for which the generated URL for recording batch download remains valid.The default duration is set to 60 minutes, with a minimum allowable duration of 2 minutes and a maximum of 60 minutes.
    RecordingBatchDownloadUrlTtl int `json:"recordingBatchDownloadUrlTtl"`

}

// String returns a JSON representation of the model
func (o *Recordingsettings) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingsettings) MarshalJSON() ([]byte, error) {
    type Alias Recordingsettings

    if RecordingsettingsMarshalled {
        return []byte("{}"), nil
    }
    RecordingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        MaxSimultaneousStreams int `json:"maxSimultaneousStreams"`
        
        MaxConfigurableScreenRecordingStreams int `json:"maxConfigurableScreenRecordingStreams"`
        
        RegionalRecordingStorageEnabled bool `json:"regionalRecordingStorageEnabled"`
        
        RecordingPlaybackUrlTtl int `json:"recordingPlaybackUrlTtl"`
        
        RecordingBatchDownloadUrlTtl int `json:"recordingBatchDownloadUrlTtl"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InitiatevideochatrecordingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InitiatevideochatrecordingDud struct { 
    


    


    

}

// Initiatevideochatrecording
type Initiatevideochatrecording struct { 
    // VideoChatRecordingCaptureType - The type of video recording to capture.
    VideoChatRecordingCaptureType string `json:"videoChatRecordingCaptureType"`


    // ArchiveRetention - Archive retention settings for the initiated video chat recording.
    ArchiveRetention Archiveretention `json:"archiveRetention"`


    // DeleteRetention - Delete retention settings for the initiated video chat recording.
    DeleteRetention Deleteretention `json:"deleteRetention"`

}

// String returns a JSON representation of the model
func (o *Initiatevideochatrecording) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Initiatevideochatrecording) MarshalJSON() ([]byte, error) {
    type Alias Initiatevideochatrecording

    if InitiatevideochatrecordingMarshalled {
        return []byte("{}"), nil
    }
    InitiatevideochatrecordingMarshalled = true

    return json.Marshal(&struct {
        
        VideoChatRecordingCaptureType string `json:"videoChatRecordingCaptureType"`
        
        ArchiveRetention Archiveretention `json:"archiveRetention"`
        
        DeleteRetention Deleteretention `json:"deleteRetention"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


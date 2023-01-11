package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordinguploadreportrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordinguploadreportrequestDud struct { 
    


    

}

// Recordinguploadreportrequest
type Recordinguploadreportrequest struct { 
    // DateSince - Report will include uploads since this date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateSince time.Time `json:"dateSince"`


    // UploadStatus - Report will include uploads with this status
    UploadStatus string `json:"uploadStatus"`

}

// String returns a JSON representation of the model
func (o *Recordinguploadreportrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordinguploadreportrequest) MarshalJSON() ([]byte, error) {
    type Alias Recordinguploadreportrequest

    if RecordinguploadreportrequestMarshalled {
        return []byte("{}"), nil
    }
    RecordinguploadreportrequestMarshalled = true

    return json.Marshal(&struct {
        
        DateSince time.Time `json:"dateSince"`
        
        UploadStatus string `json:"uploadStatus"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


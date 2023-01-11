package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordinguploadreportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordinguploadreportDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Recordinguploadreport
type Recordinguploadreport struct { 
    // Id - The report id.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // State - The current status of the upload report.
    State string `json:"state"`


    // SignedUrl - For COMPLETED tasks, the signed url to download the report.
    SignedUrl string `json:"signedUrl"`


    

}

// String returns a JSON representation of the model
func (o *Recordinguploadreport) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordinguploadreport) MarshalJSON() ([]byte, error) {
    type Alias Recordinguploadreport

    if RecordinguploadreportMarshalled {
        return []byte("{}"), nil
    }
    RecordinguploadreportMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        State string `json:"state"`
        
        SignedUrl string `json:"signedUrl"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


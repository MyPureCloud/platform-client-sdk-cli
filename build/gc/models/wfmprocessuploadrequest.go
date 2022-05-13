package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmprocessuploadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmprocessuploadrequestDud struct { 
    

}

// Wfmprocessuploadrequest
type Wfmprocessuploadrequest struct { 
    // UploadKey - The uploadKey provided by the request to get an upload URL
    UploadKey string `json:"uploadKey"`

}

// String returns a JSON representation of the model
func (o *Wfmprocessuploadrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmprocessuploadrequest) MarshalJSON() ([]byte, error) {
    type Alias Wfmprocessuploadrequest

    if WfmprocessuploadrequestMarshalled {
        return []byte("{}"), nil
    }
    WfmprocessuploadrequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


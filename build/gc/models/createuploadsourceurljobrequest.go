package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateuploadsourceurljobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateuploadsourceurljobrequestDud struct { 
    

}

// Createuploadsourceurljobrequest
type Createuploadsourceurljobrequest struct { 
    // UploadUrl - The URL of the content to upload.
    UploadUrl string `json:"uploadUrl"`

}

// String returns a JSON representation of the model
func (o *Createuploadsourceurljobrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createuploadsourceurljobrequest) MarshalJSON() ([]byte, error) {
    type Alias Createuploadsourceurljobrequest

    if CreateuploadsourceurljobrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateuploadsourceurljobrequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadUrl string `json:"uploadUrl"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


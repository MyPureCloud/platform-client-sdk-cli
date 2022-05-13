package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgelogsjobuploadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgelogsjobuploadrequestDud struct { 
    

}

// Edgelogsjobuploadrequest
type Edgelogsjobuploadrequest struct { 
    // FileIds - A list of file ids to upload.
    FileIds []string `json:"fileIds"`

}

// String returns a JSON representation of the model
func (o *Edgelogsjobuploadrequest) String() string {
     o.FileIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgelogsjobuploadrequest) MarshalJSON() ([]byte, error) {
    type Alias Edgelogsjobuploadrequest

    if EdgelogsjobuploadrequestMarshalled {
        return []byte("{}"), nil
    }
    EdgelogsjobuploadrequestMarshalled = true

    return json.Marshal(&struct {
        
        FileIds []string `json:"fileIds"`
        *Alias
    }{

        
        FileIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}


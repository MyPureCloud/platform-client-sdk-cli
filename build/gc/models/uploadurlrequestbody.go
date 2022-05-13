package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UploadurlrequestbodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UploadurlrequestbodyDud struct { 
    

}

// Uploadurlrequestbody
type Uploadurlrequestbody struct { 
    // ContentLengthBytes - The expected content length (in bytes) of the gzip-encoded data that will be PUT to the returned signed URL
    ContentLengthBytes int `json:"contentLengthBytes"`

}

// String returns a JSON representation of the model
func (o *Uploadurlrequestbody) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Uploadurlrequestbody) MarshalJSON() ([]byte, error) {
    type Alias Uploadurlrequestbody

    if UploadurlrequestbodyMarshalled {
        return []byte("{}"), nil
    }
    UploadurlrequestbodyMarshalled = true

    return json.Marshal(&struct {
        
        ContentLengthBytes int `json:"contentLengthBytes"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


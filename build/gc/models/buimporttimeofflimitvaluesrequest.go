package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuimporttimeofflimitvaluesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuimporttimeofflimitvaluesrequestDud struct { 
    

}

// Buimporttimeofflimitvaluesrequest
type Buimporttimeofflimitvaluesrequest struct { 
    // UploadKey - The uploadKey provided by the request to get an upload URL
    UploadKey string `json:"uploadKey"`

}

// String returns a JSON representation of the model
func (o *Buimporttimeofflimitvaluesrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buimporttimeofflimitvaluesrequest) MarshalJSON() ([]byte, error) {
    type Alias Buimporttimeofflimitvaluesrequest

    if BuimporttimeofflimitvaluesrequestMarshalled {
        return []byte("{}"), nil
    }
    BuimporttimeofflimitvaluesrequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


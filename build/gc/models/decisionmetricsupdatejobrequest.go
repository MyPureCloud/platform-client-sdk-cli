package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisionmetricsupdatejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisionmetricsupdatejobrequestDud struct { 
    

}

// Decisionmetricsupdatejobrequest
type Decisionmetricsupdatejobrequest struct { 
    // UploadKey - The S3 key for the uploaded decision metrics file
    UploadKey string `json:"uploadKey"`

}

// String returns a JSON representation of the model
func (o *Decisionmetricsupdatejobrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisionmetricsupdatejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Decisionmetricsupdatejobrequest

    if DecisionmetricsupdatejobrequestMarshalled {
        return []byte("{}"), nil
    }
    DecisionmetricsupdatejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


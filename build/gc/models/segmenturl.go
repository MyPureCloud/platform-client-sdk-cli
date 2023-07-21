package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmenturlMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmenturlDud struct { 
    


    

}

// Segmenturl
type Segmenturl struct { 
    // Recording - The Recording Reference
    Recording Addressableentityref `json:"recording"`


    // Url - The pre-signed S3 URL of the transcript
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Segmenturl) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmenturl) MarshalJSON() ([]byte, error) {
    type Alias Segmenturl

    if SegmenturlMarshalled {
        return []byte("{}"), nil
    }
    SegmenturlMarshalled = true

    return json.Marshal(&struct {
        
        Recording Addressableentityref `json:"recording"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


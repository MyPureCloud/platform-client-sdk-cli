package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscripturlMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscripturlDud struct { 
    

}

// Transcripturl
type Transcripturl struct { 
    // Url - The pre-signed S3 URL of the transcript
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Transcripturl) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcripturl) MarshalJSON() ([]byte, error) {
    type Alias Transcripturl

    if TranscripturlMarshalled {
        return []byte("{}"), nil
    }
    TranscripturlMarshalled = true

    return json.Marshal(&struct { 
        Url string `json:"url"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


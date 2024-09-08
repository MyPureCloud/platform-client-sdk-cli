package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptionconfigDud struct { 
    

}

// Transcriptionconfig
type Transcriptionconfig struct { 
    // VendorName - The name of the vendor used for speech transcription.
    VendorName string `json:"vendorName"`

}

// String returns a JSON representation of the model
func (o *Transcriptionconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptionconfig) MarshalJSON() ([]byte, error) {
    type Alias Transcriptionconfig

    if TranscriptionconfigMarshalled {
        return []byte("{}"), nil
    }
    TranscriptionconfigMarshalled = true

    return json.Marshal(&struct {
        
        VendorName string `json:"vendorName"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


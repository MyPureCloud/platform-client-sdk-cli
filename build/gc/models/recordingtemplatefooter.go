package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingtemplatefooterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingtemplatefooterDud struct { 
    

}

// Recordingtemplatefooter
type Recordingtemplatefooter struct { 
    // Text - Footer text.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Recordingtemplatefooter) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingtemplatefooter) MarshalJSON() ([]byte, error) {
    type Alias Recordingtemplatefooter

    if RecordingtemplatefooterMarshalled {
        return []byte("{}"), nil
    }
    RecordingtemplatefooterMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingtemplatebodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingtemplatebodyDud struct { 
    

}

// Recordingtemplatebody
type Recordingtemplatebody struct { 
    // Text - Template parameters for placeholders in template.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Recordingtemplatebody) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingtemplatebody) MarshalJSON() ([]byte, error) {
    type Alias Recordingtemplatebody

    if RecordingtemplatebodyMarshalled {
        return []byte("{}"), nil
    }
    RecordingtemplatebodyMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


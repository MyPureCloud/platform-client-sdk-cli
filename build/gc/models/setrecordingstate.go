package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetrecordingstateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetrecordingstateDud struct { 
    

}

// Setrecordingstate
type Setrecordingstate struct { 
    // RecordingState - On update, 'paused' initiates a secure pause, 'active' resumes any paused recordings.
    RecordingState string `json:"recordingState"`

}

// String returns a JSON representation of the model
func (o *Setrecordingstate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setrecordingstate) MarshalJSON() ([]byte, error) {
    type Alias Setrecordingstate

    if SetrecordingstateMarshalled {
        return []byte("{}"), nil
    }
    SetrecordingstateMarshalled = true

    return json.Marshal(&struct {
        
        RecordingState string `json:"recordingState"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


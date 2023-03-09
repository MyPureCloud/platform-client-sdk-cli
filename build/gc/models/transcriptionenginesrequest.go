package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptionenginesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptionenginesrequestDud struct { 
    

}

// Transcriptionenginesrequest
type Transcriptionenginesrequest struct { 
    // TranscriptionEngines - The transcription engine setting
    TranscriptionEngines []Transcriptionengines `json:"transcriptionEngines"`

}

// String returns a JSON representation of the model
func (o *Transcriptionenginesrequest) String() string {
     o.TranscriptionEngines = []Transcriptionengines{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptionenginesrequest) MarshalJSON() ([]byte, error) {
    type Alias Transcriptionenginesrequest

    if TranscriptionenginesrequestMarshalled {
        return []byte("{}"), nil
    }
    TranscriptionenginesrequestMarshalled = true

    return json.Marshal(&struct {
        
        TranscriptionEngines []Transcriptionengines `json:"transcriptionEngines"`
        *Alias
    }{

        
        TranscriptionEngines: []Transcriptionengines{{}},
        

        Alias: (*Alias)(u),
    })
}


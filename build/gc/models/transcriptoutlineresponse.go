package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptoutlineresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptoutlineresponseDud struct { 
    

}

// Transcriptoutlineresponse
type Transcriptoutlineresponse struct { 
    // Segments - The list of outline segments for the communication
    Segments []Transcriptoutlinesegment `json:"segments"`

}

// String returns a JSON representation of the model
func (o *Transcriptoutlineresponse) String() string {
     o.Segments = []Transcriptoutlinesegment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptoutlineresponse) MarshalJSON() ([]byte, error) {
    type Alias Transcriptoutlineresponse

    if TranscriptoutlineresponseMarshalled {
        return []byte("{}"), nil
    }
    TranscriptoutlineresponseMarshalled = true

    return json.Marshal(&struct {
        
        Segments []Transcriptoutlinesegment `json:"segments"`
        *Alias
    }{

        
        Segments: []Transcriptoutlinesegment{{}},
        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptsDud struct { 
    


    


    

}

// Transcripts
type Transcripts struct { 
    // ExactMatch - List of transcript contents which needs to satisfy exact match criteria
    ExactMatch []string `json:"exactMatch"`


    // Contains - List of transcript contents which needs to satisfy contains criteria
    Contains []string `json:"contains"`


    // DoesNotContain - List of transcript contents which needs to satisfy does not contain criteria
    DoesNotContain []string `json:"doesNotContain"`

}

// String returns a JSON representation of the model
func (o *Transcripts) String() string {
     o.ExactMatch = []string{""} 
     o.Contains = []string{""} 
     o.DoesNotContain = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcripts) MarshalJSON() ([]byte, error) {
    type Alias Transcripts

    if TranscriptsMarshalled {
        return []byte("{}"), nil
    }
    TranscriptsMarshalled = true

    return json.Marshal(&struct {
        
        ExactMatch []string `json:"exactMatch"`
        
        Contains []string `json:"contains"`
        
        DoesNotContain []string `json:"doesNotContain"`
        *Alias
    }{

        
        ExactMatch: []string{""},
        


        
        Contains: []string{""},
        


        
        DoesNotContain: []string{""},
        

        Alias: (*Alias)(u),
    })
}


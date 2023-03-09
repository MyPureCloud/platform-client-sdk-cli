package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptionenginesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptionenginesDud struct { 
    


    

}

// Transcriptionengines
type Transcriptionengines struct { 
    // Engine
    Engine string `json:"engine"`


    // Dialects
    Dialects []string `json:"dialects"`

}

// String returns a JSON representation of the model
func (o *Transcriptionengines) String() string {
    
     o.Dialects = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptionengines) MarshalJSON() ([]byte, error) {
    type Alias Transcriptionengines

    if TranscriptionenginesMarshalled {
        return []byte("{}"), nil
    }
    TranscriptionenginesMarshalled = true

    return json.Marshal(&struct {
        
        Engine string `json:"engine"`
        
        Dialects []string `json:"dialects"`
        *Alias
    }{

        


        
        Dialects: []string{""},
        

        Alias: (*Alias)(u),
    })
}


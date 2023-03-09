package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupporteddialectsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupporteddialectsentitylistingDud struct { 
    

}

// Supporteddialectsentitylisting
type Supporteddialectsentitylisting struct { 
    // Entities
    Entities []Transcriptionengines `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Supporteddialectsentitylisting) String() string {
     o.Entities = []Transcriptionengines{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supporteddialectsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Supporteddialectsentitylisting

    if SupporteddialectsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SupporteddialectsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Transcriptionengines `json:"entities"`
        *Alias
    }{

        
        Entities: []Transcriptionengines{{}},
        

        Alias: (*Alias)(u),
    })
}


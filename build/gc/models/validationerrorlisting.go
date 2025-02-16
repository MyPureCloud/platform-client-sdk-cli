package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidationerrorlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidationerrorlistingDud struct { 
    

}

// Validationerrorlisting
type Validationerrorlisting struct { 
    // Entities
    Entities []Validationerrorresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Validationerrorlisting) String() string {
     o.Entities = []Validationerrorresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validationerrorlisting) MarshalJSON() ([]byte, error) {
    type Alias Validationerrorlisting

    if ValidationerrorlistingMarshalled {
        return []byte("{}"), nil
    }
    ValidationerrorlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Validationerrorresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Validationerrorresponse{{}},
        

        Alias: (*Alias)(u),
    })
}


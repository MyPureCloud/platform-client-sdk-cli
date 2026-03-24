package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntakesettingslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntakesettingslistingDud struct { 
    

}

// Intakesettingslisting
type Intakesettingslisting struct { 
    // Entities
    Entities []Intakesetting `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Intakesettingslisting) String() string {
     o.Entities = []Intakesetting{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intakesettingslisting) MarshalJSON() ([]byte, error) {
    type Alias Intakesettingslisting

    if IntakesettingslistingMarshalled {
        return []byte("{}"), nil
    }
    IntakesettingslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Intakesetting `json:"entities"`
        *Alias
    }{

        
        Entities: []Intakesetting{{}},
        

        Alias: (*Alias)(u),
    })
}


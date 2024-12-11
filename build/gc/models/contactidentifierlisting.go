package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactidentifierlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactidentifierlistingDud struct { 
    

}

// Contactidentifierlisting
type Contactidentifierlisting struct { 
    // Entities
    Entities []Contactidentifier `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Contactidentifierlisting) String() string {
     o.Entities = []Contactidentifier{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactidentifierlisting) MarshalJSON() ([]byte, error) {
    type Alias Contactidentifierlisting

    if ContactidentifierlistingMarshalled {
        return []byte("{}"), nil
    }
    ContactidentifierlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contactidentifier `json:"entities"`
        *Alias
    }{

        
        Entities: []Contactidentifier{{}},
        

        Alias: (*Alias)(u),
    })
}


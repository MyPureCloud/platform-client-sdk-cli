package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalorganizationidentifierlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalorganizationidentifierlistingDud struct { 
    

}

// Externalorganizationidentifierlisting
type Externalorganizationidentifierlisting struct { 
    // Entities
    Entities []Externalorganizationidentifier `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Externalorganizationidentifierlisting) String() string {
     o.Entities = []Externalorganizationidentifier{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalorganizationidentifierlisting) MarshalJSON() ([]byte, error) {
    type Alias Externalorganizationidentifierlisting

    if ExternalorganizationidentifierlistingMarshalled {
        return []byte("{}"), nil
    }
    ExternalorganizationidentifierlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Externalorganizationidentifier `json:"entities"`
        *Alias
    }{

        
        Entities: []Externalorganizationidentifier{{}},
        

        Alias: (*Alias)(u),
    })
}


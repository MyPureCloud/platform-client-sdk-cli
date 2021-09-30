package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SecuresessionentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SecuresessionentitylistingDud struct { 
    

}

// Securesessionentitylisting
type Securesessionentitylisting struct { 
    // Entities
    Entities []Securesession `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Securesessionentitylisting) String() string {
    
    
     o.Entities = []Securesession{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Securesessionentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Securesessionentitylisting

    if SecuresessionentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SecuresessionentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Securesession `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Securesession{{}},
        

        
        Alias: (*Alias)(u),
    })
}


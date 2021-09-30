package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LimitsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LimitsentitylistingDud struct { 
    

}

// Limitsentitylisting
type Limitsentitylisting struct { 
    // Entities
    Entities []Limit `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Limitsentitylisting) String() string {
    
    
     o.Entities = []Limit{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Limitsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Limitsentitylisting

    if LimitsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    LimitsentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Limit `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Limit{{}},
        

        
        Alias: (*Alias)(u),
    })
}


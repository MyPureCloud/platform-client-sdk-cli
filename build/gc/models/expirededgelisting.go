package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExpirededgelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExpirededgelistingDud struct { 
    

}

// Expirededgelisting
type Expirededgelisting struct { 
    // Entities
    Entities []Domainentityref `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Expirededgelisting) String() string {
    
    
     o.Entities = []Domainentityref{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Expirededgelisting) MarshalJSON() ([]byte, error) {
    type Alias Expirededgelisting

    if ExpirededgelistingMarshalled {
        return []byte("{}"), nil
    }
    ExpirededgelistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Domainentityref `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Domainentityref{{}},
        

        
        Alias: (*Alias)(u),
    })
}


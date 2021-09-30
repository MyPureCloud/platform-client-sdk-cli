package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulelistingDud struct { 
    

}

// Buschedulelisting
type Buschedulelisting struct { 
    // Entities
    Entities []Buschedulelistitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Buschedulelisting) String() string {
    
    
     o.Entities = []Buschedulelistitem{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulelisting) MarshalJSON() ([]byte, error) {
    type Alias Buschedulelisting

    if BuschedulelistingMarshalled {
        return []byte("{}"), nil
    }
    BuschedulelistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Buschedulelistitem `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Buschedulelistitem{{}},
        

        
        Alias: (*Alias)(u),
    })
}


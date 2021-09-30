package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EntitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EntitylistingDud struct { 
    

}

// Entitylisting
type Entitylisting struct { 
    // Entities
    Entities []interface{} `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Entitylisting) String() string {
    
    
     o.Entities = []interface{}{} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Entitylisting) MarshalJSON() ([]byte, error) {
    type Alias Entitylisting

    if EntitylistingMarshalled {
        return []byte("{}"), nil
    }
    EntitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []interface{} `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []interface{}{},
        

        
        Alias: (*Alias)(u),
    })
}


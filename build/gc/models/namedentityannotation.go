package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NamedentityannotationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NamedentityannotationDud struct { 
    

}

// Namedentityannotation
type Namedentityannotation struct { 
    // Name - The name of the annotated named entity.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Namedentityannotation) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Namedentityannotation) MarshalJSON() ([]byte, error) {
    type Alias Namedentityannotation

    if NamedentityannotationMarshalled {
        return []byte("{}"), nil
    }
    NamedentityannotationMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


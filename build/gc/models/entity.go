package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EntityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EntityDud struct { 
    Id string `json:"id"`

}

// Entity
type Entity struct { 
    

}

// String returns a JSON representation of the model
func (o *Entity) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Entity) MarshalJSON() ([]byte, error) {
    type Alias Entity

    if EntityMarshalled {
        return []byte("{}"), nil
    }
    EntityMarshalled = true

    return json.Marshal(&struct { 
        
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


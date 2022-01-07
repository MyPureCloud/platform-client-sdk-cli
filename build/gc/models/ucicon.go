package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UciconMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UciconDud struct { 
    Vector string `json:"vector"`

}

// Ucicon
type Ucicon struct { 
    

}

// String returns a JSON representation of the model
func (o *Ucicon) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ucicon) MarshalJSON() ([]byte, error) {
    type Alias Ucicon

    if UciconMarshalled {
        return []byte("{}"), nil
    }
    UciconMarshalled = true

    return json.Marshal(&struct { 
        
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


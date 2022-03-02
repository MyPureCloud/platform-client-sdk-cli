package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicdurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicdurationDud struct { 
    TotalMilliseconds int `json:"totalMilliseconds"`

}

// Topicduration
type Topicduration struct { 
    

}

// String returns a JSON representation of the model
func (o *Topicduration) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicduration) MarshalJSON() ([]byte, error) {
    type Alias Topicduration

    if TopicdurationMarshalled {
        return []byte("{}"), nil
    }
    TopicdurationMarshalled = true

    return json.Marshal(&struct { 
        
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


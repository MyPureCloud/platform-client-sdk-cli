package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContextintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContextintentDud struct { 
    

}

// Contextintent
type Contextintent struct { 
    // Name - The name of the intent.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Contextintent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contextintent) MarshalJSON() ([]byte, error) {
    type Alias Contextintent

    if ContextintentMarshalled {
        return []byte("{}"), nil
    }
    ContextintentMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


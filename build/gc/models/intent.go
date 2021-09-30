package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntentDud struct { 
    

}

// Intent
type Intent struct { 
    // Name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Intent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intent) MarshalJSON() ([]byte, error) {
    type Alias Intent

    if IntentMarshalled {
        return []byte("{}"), nil
    }
    IntentMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


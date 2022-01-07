package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Uci10nMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Uci10nDud struct { 
    

}

// Uci10n
type Uci10n struct { 
    // Name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Uci10n) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Uci10n) MarshalJSON() ([]byte, error) {
    type Alias Uci10n

    if Uci10nMarshalled {
        return []byte("{}"), nil
    }
    Uci10nMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessengerhomescreenMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessengerhomescreenDud struct { 
    

}

// Messengerhomescreen
type Messengerhomescreen struct { 
    // Enabled - whether or not homescreen is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Messengerhomescreen) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messengerhomescreen) MarshalJSON() ([]byte, error) {
    type Alias Messengerhomescreen

    if MessengerhomescreenMarshalled {
        return []byte("{}"), nil
    }
    MessengerhomescreenMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


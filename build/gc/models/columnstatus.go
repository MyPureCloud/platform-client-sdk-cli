package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ColumnstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ColumnstatusDud struct { 
    

}

// Columnstatus
type Columnstatus struct { 
    // Contactable - Indicates whether or not an individual contact method column is contactable.
    Contactable bool `json:"contactable"`

}

// String returns a JSON representation of the model
func (o *Columnstatus) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Columnstatus) MarshalJSON() ([]byte, error) {
    type Alias Columnstatus

    if ColumnstatusMarshalled {
        return []byte("{}"), nil
    }
    ColumnstatusMarshalled = true

    return json.Marshal(&struct { 
        Contactable bool `json:"contactable"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


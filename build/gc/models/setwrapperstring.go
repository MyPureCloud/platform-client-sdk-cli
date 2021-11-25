package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetwrapperstringMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetwrapperstringDud struct { 
    

}

// Setwrapperstring
type Setwrapperstring struct { 
    // Values
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Setwrapperstring) String() string {
    
    
     o.Values = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setwrapperstring) MarshalJSON() ([]byte, error) {
    type Alias Setwrapperstring

    if SetwrapperstringMarshalled {
        return []byte("{}"), nil
    }
    SetwrapperstringMarshalled = true

    return json.Marshal(&struct { 
        Values []string `json:"values"`
        
        *Alias
    }{
        

        
        Values: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}


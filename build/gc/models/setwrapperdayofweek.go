package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetwrapperdayofweekMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetwrapperdayofweekDud struct { 
    

}

// Setwrapperdayofweek
type Setwrapperdayofweek struct { 
    // Values
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Setwrapperdayofweek) String() string {
    
    
     o.Values = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setwrapperdayofweek) MarshalJSON() ([]byte, error) {
    type Alias Setwrapperdayofweek

    if SetwrapperdayofweekMarshalled {
        return []byte("{}"), nil
    }
    SetwrapperdayofweekMarshalled = true

    return json.Marshal(&struct { 
        Values []string `json:"values"`
        
        *Alias
    }{
        

        
        Values: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}


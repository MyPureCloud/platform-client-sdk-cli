package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AccessrestrictionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AccessrestrictionDud struct { 
    


    

}

// Accessrestriction
type Accessrestriction struct { 
    // Kind
    Kind string `json:"kind"`


    // Values
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Accessrestriction) String() string {
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Accessrestriction) MarshalJSON() ([]byte, error) {
    type Alias Accessrestriction

    if AccessrestrictionMarshalled {
        return []byte("{}"), nil
    }
    AccessrestrictionMarshalled = true

    return json.Marshal(&struct {
        
        Kind string `json:"kind"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}


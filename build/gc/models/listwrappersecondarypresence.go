package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrappersecondarypresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrappersecondarypresenceDud struct { 
    

}

// Listwrappersecondarypresence
type Listwrappersecondarypresence struct { 
    // Values
    Values []Secondarypresence `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrappersecondarypresence) String() string {
    
    
     o.Values = []Secondarypresence{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrappersecondarypresence) MarshalJSON() ([]byte, error) {
    type Alias Listwrappersecondarypresence

    if ListwrappersecondarypresenceMarshalled {
        return []byte("{}"), nil
    }
    ListwrappersecondarypresenceMarshalled = true

    return json.Marshal(&struct { 
        Values []Secondarypresence `json:"values"`
        
        *Alias
    }{
        

        
        Values: []Secondarypresence{{}},
        

        
        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MemberentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MemberentityDud struct { 
    

}

// Memberentity
type Memberentity struct { 
    // Id
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Memberentity) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Memberentity) MarshalJSON() ([]byte, error) {
    type Alias Memberentity

    if MemberentityMarshalled {
        return []byte("{}"), nil
    }
    MemberentityMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}


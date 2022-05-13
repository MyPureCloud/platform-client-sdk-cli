package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalsegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalsegmentDud struct { 
    


    


    

}

// Externalsegment
type Externalsegment struct { 
    // Id - Identifier for the external segment in the system where it originates from.
    Id string `json:"id"`


    // Name - Name for the external segment in the system where it originates from.
    Name string `json:"name"`


    // Source - The external system where the segment originates from.
    Source string `json:"source"`

}

// String returns a JSON representation of the model
func (o *Externalsegment) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalsegment) MarshalJSON() ([]byte, error) {
    type Alias Externalsegment

    if ExternalsegmentMarshalled {
        return []byte("{}"), nil
    }
    ExternalsegmentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Source string `json:"source"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


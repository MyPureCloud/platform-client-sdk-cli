package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NamedentitytypemechanismexampleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NamedentitytypemechanismexampleDud struct { 
    


    

}

// Namedentitytypemechanismexample
type Namedentitytypemechanismexample struct { 
    // Text - Example input text
    Text string `json:"text"`


    // ResolvedValue - Resolved entity value
    ResolvedValue string `json:"resolvedValue"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypemechanismexample) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Namedentitytypemechanismexample) MarshalJSON() ([]byte, error) {
    type Alias Namedentitytypemechanismexample

    if NamedentitytypemechanismexampleMarshalled {
        return []byte("{}"), nil
    }
    NamedentitytypemechanismexampleMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        ResolvedValue string `json:"resolvedValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


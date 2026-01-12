package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnavailabletimesviolationmessageargumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnavailabletimesviolationmessageargumentDud struct { 
    


    

}

// Unavailabletimesviolationmessageargument
type Unavailabletimesviolationmessageargument struct { 
    // VarType - Argument type
    VarType string `json:"type"`


    // Value - Value of argument
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Unavailabletimesviolationmessageargument) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unavailabletimesviolationmessageargument) MarshalJSON() ([]byte, error) {
    type Alias Unavailabletimesviolationmessageargument

    if UnavailabletimesviolationmessageargumentMarshalled {
        return []byte("{}"), nil
    }
    UnavailabletimesviolationmessageargumentMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


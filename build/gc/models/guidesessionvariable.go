package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionvariableMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionvariableDud struct { 
    


    

}

// Guidesessionvariable
type Guidesessionvariable struct { 
    // Name - The name of the variable.
    Name string `json:"name"`


    // Value - The value of the variable.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Guidesessionvariable) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionvariable) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionvariable

    if GuidesessionvariableMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionvariableMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


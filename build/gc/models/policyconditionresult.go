package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicyconditionresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicyconditionresultDud struct { 
    


    

}

// Policyconditionresult
type Policyconditionresult struct { 
    // Name - The condition name
    Name string `json:"name"`


    // Result - The boolean result of the condition
    Result bool `json:"result"`

}

// String returns a JSON representation of the model
func (o *Policyconditionresult) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policyconditionresult) MarshalJSON() ([]byte, error) {
    type Alias Policyconditionresult

    if PolicyconditionresultMarshalled {
        return []byte("{}"), nil
    }
    PolicyconditionresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Result bool `json:"result"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


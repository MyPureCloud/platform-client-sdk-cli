package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MatchtestresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MatchtestresultDud struct { 
    


    

}

// Matchtestresult - Information about the results being matched by the expressions
type Matchtestresult struct { 
    // Value - The value of the field being matched
    Value interface{} `json:"value"`


    // Path - The json path to the json node being matched on. ex: $['things'][1]
    Path string `json:"path"`

}

// String returns a JSON representation of the model
func (o *Matchtestresult) String() string {
     o.Value = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Matchtestresult) MarshalJSON() ([]byte, error) {
    type Alias Matchtestresult

    if MatchtestresultMarshalled {
        return []byte("{}"), nil
    }
    MatchtestresultMarshalled = true

    return json.Marshal(&struct {
        
        Value interface{} `json:"value"`
        
        Path string `json:"path"`
        *Alias
    }{

        
        Value: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}


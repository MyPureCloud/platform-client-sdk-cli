package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TypedattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TypedattributeDud struct { 
    


    

}

// Typedattribute
type Typedattribute struct { 
    // VarType
    VarType string `json:"type"`


    // Value
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Typedattribute) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Typedattribute) MarshalJSON() ([]byte, error) {
    type Alias Typedattribute

    if TypedattributeMarshalled {
        return []byte("{}"), nil
    }
    TypedattributeMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


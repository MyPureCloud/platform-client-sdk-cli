package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TargetattributevalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TargetattributevalueDud struct { 
    


    

}

// Targetattributevalue
type Targetattributevalue struct { 
    // Description
    Description string `json:"description"`


    // PolicyAttributes
    PolicyAttributes []Policyattribute `json:"policyAttributes"`

}

// String returns a JSON representation of the model
func (o *Targetattributevalue) String() string {
    
     o.PolicyAttributes = []Policyattribute{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Targetattributevalue) MarshalJSON() ([]byte, error) {
    type Alias Targetattributevalue

    if TargetattributevalueMarshalled {
        return []byte("{}"), nil
    }
    TargetattributevalueMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        PolicyAttributes []Policyattribute `json:"policyAttributes"`
        *Alias
    }{

        


        
        PolicyAttributes: []Policyattribute{{}},
        

        Alias: (*Alias)(u),
    })
}


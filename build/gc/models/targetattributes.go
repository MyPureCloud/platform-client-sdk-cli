package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TargetattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TargetattributesDud struct { 
    


    

}

// Targetattributes
type Targetattributes struct { 
    // BaseAttributes - A set of base attributes which may be used in policies for any target.
    BaseAttributes []Policyattribute `json:"baseAttributes"`


    // TargetAttributes - A map of policy targets to any additional attributes which are valid for that target.
    TargetAttributes map[string]Targetattributevalue `json:"targetAttributes"`

}

// String returns a JSON representation of the model
func (o *Targetattributes) String() string {
     o.BaseAttributes = []Policyattribute{{}} 
     o.TargetAttributes = map[string]Targetattributevalue{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Targetattributes) MarshalJSON() ([]byte, error) {
    type Alias Targetattributes

    if TargetattributesMarshalled {
        return []byte("{}"), nil
    }
    TargetattributesMarshalled = true

    return json.Marshal(&struct {
        
        BaseAttributes []Policyattribute `json:"baseAttributes"`
        
        TargetAttributes map[string]Targetattributevalue `json:"targetAttributes"`
        *Alias
    }{

        
        BaseAttributes: []Policyattribute{{}},
        


        
        TargetAttributes: map[string]Targetattributevalue{"": {}},
        

        Alias: (*Alias)(u),
    })
}


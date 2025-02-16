package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicyattributesetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicyattributesetDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Policyattributeset
type Policyattributeset struct { 
    


    // Name
    Name string `json:"name"`


    // PolicyAttributes - A set of the attributes checked by the requested policy.
    PolicyAttributes []Policyattribute `json:"policyAttributes"`


    // PresetAttributes - Map of names and values of preset attributes used in this policy.
    PresetAttributes map[string]Typedattribute `json:"presetAttributes"`


    

}

// String returns a JSON representation of the model
func (o *Policyattributeset) String() string {
    
     o.PolicyAttributes = []Policyattribute{{}} 
     o.PresetAttributes = map[string]Typedattribute{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policyattributeset) MarshalJSON() ([]byte, error) {
    type Alias Policyattributeset

    if PolicyattributesetMarshalled {
        return []byte("{}"), nil
    }
    PolicyattributesetMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        PolicyAttributes []Policyattribute `json:"policyAttributes"`
        
        PresetAttributes map[string]Typedattribute `json:"presetAttributes"`
        *Alias
    }{

        


        


        
        PolicyAttributes: []Policyattribute{{}},
        


        
        PresetAttributes: map[string]Typedattribute{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}


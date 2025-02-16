package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicyattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicyattributeDud struct { 
    


    


    


    

}

// Policyattribute
type Policyattribute struct { 
    // Name
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`


    // Description
    Description string `json:"description"`


    // FeatureToggle
    FeatureToggle string `json:"featureToggle"`

}

// String returns a JSON representation of the model
func (o *Policyattribute) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policyattribute) MarshalJSON() ([]byte, error) {
    type Alias Policyattribute

    if PolicyattributeMarshalled {
        return []byte("{}"), nil
    }
    PolicyattributeMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Description string `json:"description"`
        
        FeatureToggle string `json:"featureToggle"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


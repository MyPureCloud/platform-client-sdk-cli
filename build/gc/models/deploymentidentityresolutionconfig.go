package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DeploymentidentityresolutionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DeploymentidentityresolutionconfigDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Deploymentidentityresolutionconfig
type Deploymentidentityresolutionconfig struct { 
    


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    

}

// String returns a JSON representation of the model
func (o *Deploymentidentityresolutionconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Deploymentidentityresolutionconfig) MarshalJSON() ([]byte, error) {
    type Alias Deploymentidentityresolutionconfig

    if DeploymentidentityresolutionconfigMarshalled {
        return []byte("{}"), nil
    }
    DeploymentidentityresolutionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        ResolveIdentities bool `json:"resolveIdentities"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


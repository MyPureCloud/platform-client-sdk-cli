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
    


    // Division - The division to use when performing identity resolution.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    // ExternalSource - The external source used for stitching this channel.
    ExternalSource Identityresolutionexternalsource `json:"externalSource"`


    // Automerge - Whether automerging of contacts should be enabled for each channel.
    Automerge Identityresolutionautomergeconfig `json:"automerge"`


    

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
        
        ExternalSource Identityresolutionexternalsource `json:"externalSource"`
        
        Automerge Identityresolutionautomergeconfig `json:"automerge"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


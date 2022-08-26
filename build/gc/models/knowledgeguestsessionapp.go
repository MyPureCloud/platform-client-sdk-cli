package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestsessionappMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestsessionappDud struct { 
    


    

}

// Knowledgeguestsessionapp
type Knowledgeguestsessionapp struct { 
    // DeploymentId - App deployment ID.
    DeploymentId string `json:"deploymentId"`


    // VarType - App type.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestsessionapp) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestsessionapp) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestsessionapp

    if KnowledgeguestsessionappMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestsessionappMarshalled = true

    return json.Marshal(&struct {
        
        DeploymentId string `json:"deploymentId"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


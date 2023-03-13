package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TriggertargetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TriggertargetDud struct { 
    


    


    

}

// Triggertarget - The target of a trigger invocation
type Triggertarget struct { 
    // VarType - The entity type to target
    VarType string `json:"type"`


    // Id - The ID of the entity to target
    Id string `json:"id"`


    // WorkflowTargetSettings - Optional config for the target. Until the feature gets enabled will always operate in TopLevelPrimitives mode.
    WorkflowTargetSettings Workflowtargetsettings `json:"workflowTargetSettings"`

}

// String returns a JSON representation of the model
func (o *Triggertarget) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Triggertarget) MarshalJSON() ([]byte, error) {
    type Alias Triggertarget

    if TriggertargetMarshalled {
        return []byte("{}"), nil
    }
    TriggertargetMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Id string `json:"id"`
        
        WorkflowTargetSettings Workflowtargetsettings `json:"workflowTargetSettings"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


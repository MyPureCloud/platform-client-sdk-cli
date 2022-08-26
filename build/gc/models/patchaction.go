package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchactionDud struct { 
    


    


    


    


    


    


    


    

}

// Patchaction
type Patchaction struct { 
    // MediaType - Media type of action.
    MediaType string `json:"mediaType"`


    // ActionTemplate - Action template associated with the action map.
    ActionTemplate Actionmapactiontemplate `json:"actionTemplate"`


    // ActionTargetId - Action target ID.
    ActionTargetId string `json:"actionTargetId"`


    // IsPacingEnabled - Whether this action should be throttled.
    IsPacingEnabled bool `json:"isPacingEnabled"`


    // Props - Additional properties.
    Props Patchactionproperties `json:"props"`


    // ArchitectFlowFields - Architect Flow Id and input contract.
    ArchitectFlowFields Architectflowfields `json:"architectFlowFields"`


    // WebMessagingOfferFields - Admin-configurable fields of a web messaging offer action.
    WebMessagingOfferFields Patchwebmessagingofferfields `json:"webMessagingOfferFields"`


    // OpenActionFields - Admin-configurable fields of an open action.
    OpenActionFields Openactionfields `json:"openActionFields"`

}

// String returns a JSON representation of the model
func (o *Patchaction) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchaction) MarshalJSON() ([]byte, error) {
    type Alias Patchaction

    if PatchactionMarshalled {
        return []byte("{}"), nil
    }
    PatchactionMarshalled = true

    return json.Marshal(&struct {
        
        MediaType string `json:"mediaType"`
        
        ActionTemplate Actionmapactiontemplate `json:"actionTemplate"`
        
        ActionTargetId string `json:"actionTargetId"`
        
        IsPacingEnabled bool `json:"isPacingEnabled"`
        
        Props Patchactionproperties `json:"props"`
        
        ArchitectFlowFields Architectflowfields `json:"architectFlowFields"`
        
        WebMessagingOfferFields Patchwebmessagingofferfields `json:"webMessagingOfferFields"`
        
        OpenActionFields Openactionfields `json:"openActionFields"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


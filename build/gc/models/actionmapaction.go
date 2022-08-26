package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmapactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmapactionDud struct { 
    


    


    


    


    


    


    


    

}

// Actionmapaction
type Actionmapaction struct { 
    // ActionTemplate - Action template associated with the action map.
    ActionTemplate Actionmapactiontemplate `json:"actionTemplate"`


    // MediaType - Media type of action.
    MediaType string `json:"mediaType"`


    // ActionTargetId - Action target ID.
    ActionTargetId string `json:"actionTargetId"`


    // IsPacingEnabled - Whether this action should be throttled.
    IsPacingEnabled bool `json:"isPacingEnabled"`


    // Props - Additional properties.
    Props Actionproperties `json:"props"`


    // ArchitectFlowFields - Architect Flow Id and input contract.
    ArchitectFlowFields Architectflowfields `json:"architectFlowFields"`


    // WebMessagingOfferFields - Admin-configurable fields of a web messaging offer action.
    WebMessagingOfferFields Webmessagingofferfields `json:"webMessagingOfferFields"`


    // OpenActionFields - Admin-configurable fields of an open action.
    OpenActionFields Openactionfields `json:"openActionFields"`

}

// String returns a JSON representation of the model
func (o *Actionmapaction) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmapaction) MarshalJSON() ([]byte, error) {
    type Alias Actionmapaction

    if ActionmapactionMarshalled {
        return []byte("{}"), nil
    }
    ActionmapactionMarshalled = true

    return json.Marshal(&struct {
        
        ActionTemplate Actionmapactiontemplate `json:"actionTemplate"`
        
        MediaType string `json:"mediaType"`
        
        ActionTargetId string `json:"actionTargetId"`
        
        IsPacingEnabled bool `json:"isPacingEnabled"`
        
        Props Actionproperties `json:"props"`
        
        ArchitectFlowFields Architectflowfields `json:"architectFlowFields"`
        
        WebMessagingOfferFields Webmessagingofferfields `json:"webMessagingOfferFields"`
        
        OpenActionFields Openactionfields `json:"openActionFields"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


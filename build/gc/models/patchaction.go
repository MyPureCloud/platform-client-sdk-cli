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


    // ArchitectFlowFields - Architect Flow Id and input contract.
    ArchitectFlowFields Architectflowfields `json:"architectFlowFields"`


    // WebMessagingOfferFields - Admin-configurable fields of a web messaging offer action.
    WebMessagingOfferFields Webmessagingofferfields `json:"webMessagingOfferFields"`

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
        
        ArchitectFlowFields Architectflowfields `json:"architectFlowFields"`
        
        WebMessagingOfferFields Webmessagingofferfields `json:"webMessagingOfferFields"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


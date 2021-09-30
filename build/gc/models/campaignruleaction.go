package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleactionDud struct { 
    


    


    


    

}

// Campaignruleaction
type Campaignruleaction struct { 
    // Id
    Id string `json:"id"`


    // Parameters - The parameters for the CampaignRuleAction. Required for certain actionTypes.
    Parameters Campaignruleparameters `json:"parameters"`


    // ActionType - The action to take on the campaignRuleActionEntities.
    ActionType string `json:"actionType"`


    // CampaignRuleActionEntities - The list of entities that this action will apply to.
    CampaignRuleActionEntities Campaignruleactionentities `json:"campaignRuleActionEntities"`

}

// String returns a JSON representation of the model
func (o *Campaignruleaction) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleaction) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleaction

    if CampaignruleactionMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleactionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Parameters Campaignruleparameters `json:"parameters"`
        
        ActionType string `json:"actionType"`
        
        CampaignRuleActionEntities Campaignruleactionentities `json:"campaignRuleActionEntities"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


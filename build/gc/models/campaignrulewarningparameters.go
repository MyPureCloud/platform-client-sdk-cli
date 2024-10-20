package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignrulewarningparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignrulewarningparametersDud struct { 
    


    


    


    

}

// Campaignrulewarningparameters
type Campaignrulewarningparameters struct { 
    // ActionId - ID of action
    ActionId string `json:"actionId"`


    // ConditionId - ID of condition
    ConditionId string `json:"conditionId"`


    // ActionType - Type of action
    ActionType string `json:"actionType"`


    // ConditionType - Type of condition
    ConditionType string `json:"conditionType"`

}

// String returns a JSON representation of the model
func (o *Campaignrulewarningparameters) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignrulewarningparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignrulewarningparameters

    if CampaignrulewarningparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignrulewarningparametersMarshalled = true

    return json.Marshal(&struct {
        
        ActionId string `json:"actionId"`
        
        ConditionId string `json:"conditionId"`
        
        ActionType string `json:"actionType"`
        
        ConditionType string `json:"conditionType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


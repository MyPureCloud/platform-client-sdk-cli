package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleconditionDud struct { 
    


    


    

}

// Campaignrulecondition
type Campaignrulecondition struct { 
    // Id
    Id string `json:"id"`


    // Parameters - The parameters for the CampaignRuleCondition.
    Parameters Campaignruleparameters `json:"parameters"`


    // ConditionType - The type of condition to evaluate.
    ConditionType string `json:"conditionType"`

}

// String returns a JSON representation of the model
func (o *Campaignrulecondition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignrulecondition) MarshalJSON() ([]byte, error) {
    type Alias Campaignrulecondition

    if CampaignruleconditionMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleconditionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Parameters Campaignruleparameters `json:"parameters"`
        
        ConditionType string `json:"conditionType"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


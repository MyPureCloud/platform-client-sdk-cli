package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleconditiongroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleconditiongroupDud struct { 
    


    

}

// Campaignruleconditiongroup
type Campaignruleconditiongroup struct { 
    // MatchAnyConditions - Whether or not this condition group should be evaluated as true if any of sub conditions is matched
    MatchAnyConditions bool `json:"matchAnyConditions"`


    // Conditions - The parameters for the CampaignRuleCondition.
    Conditions []Campaignrulecondition `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Campaignruleconditiongroup) String() string {
    
     o.Conditions = []Campaignrulecondition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleconditiongroup) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleconditiongroup

    if CampaignruleconditiongroupMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleconditiongroupMarshalled = true

    return json.Marshal(&struct {
        
        MatchAnyConditions bool `json:"matchAnyConditions"`
        
        Conditions []Campaignrulecondition `json:"conditions"`
        *Alias
    }{

        


        
        Conditions: []Campaignrulecondition{{}},
        

        Alias: (*Alias)(u),
    })
}


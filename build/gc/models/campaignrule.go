package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Campaignrule
type Campaignrule struct { 
    


    // Name - The name of the CampaignRule.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // CampaignRuleEntities - The list of entities that this CampaignRule monitors.
    CampaignRuleEntities Campaignruleentities `json:"campaignRuleEntities"`


    // CampaignRuleConditions - The list of conditions that are evaluated on the entities.
    CampaignRuleConditions []Campaignrulecondition `json:"campaignRuleConditions"`


    // CampaignRuleActions - The list of actions that are executed if the conditions are satisfied.
    CampaignRuleActions []Campaignruleaction `json:"campaignRuleActions"`


    // MatchAnyConditions
    MatchAnyConditions bool `json:"matchAnyConditions"`


    // Enabled - Whether or not this CampaignRule is currently enabled. Required on updates.
    Enabled bool `json:"enabled"`


    // CampaignRuleProcessing - CampaignRule processing algorithm
    CampaignRuleProcessing string `json:"campaignRuleProcessing"`


    // ConditionGroups - List of condition groups that are evaluated, used only with campaignRuleProcessing=\"v2\"
    ConditionGroups []Campaignruleconditiongroup `json:"conditionGroups"`


    // ExecutionSettings - CampaignRule execution settings
    ExecutionSettings Campaignruleexecutionsettings `json:"executionSettings"`


    

}

// String returns a JSON representation of the model
func (o *Campaignrule) String() string {
    
    
    
     o.CampaignRuleConditions = []Campaignrulecondition{{}} 
     o.CampaignRuleActions = []Campaignruleaction{{}} 
    
    
    
     o.ConditionGroups = []Campaignruleconditiongroup{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignrule) MarshalJSON() ([]byte, error) {
    type Alias Campaignrule

    if CampaignruleMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        CampaignRuleEntities Campaignruleentities `json:"campaignRuleEntities"`
        
        CampaignRuleConditions []Campaignrulecondition `json:"campaignRuleConditions"`
        
        CampaignRuleActions []Campaignruleaction `json:"campaignRuleActions"`
        
        MatchAnyConditions bool `json:"matchAnyConditions"`
        
        Enabled bool `json:"enabled"`
        
        CampaignRuleProcessing string `json:"campaignRuleProcessing"`
        
        ConditionGroups []Campaignruleconditiongroup `json:"conditionGroups"`
        
        ExecutionSettings Campaignruleexecutionsettings `json:"executionSettings"`
        *Alias
    }{

        


        


        


        


        


        


        
        CampaignRuleConditions: []Campaignrulecondition{{}},
        


        
        CampaignRuleActions: []Campaignruleaction{{}},
        


        


        


        


        
        ConditionGroups: []Campaignruleconditiongroup{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}


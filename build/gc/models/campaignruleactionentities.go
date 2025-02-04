package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleactionentitiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleactionentitiesDud struct { 
    


    


    


    


    

}

// Campaignruleactionentities
type Campaignruleactionentities struct { 
    // Campaigns - The list of campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a campaign.
    Campaigns []Domainentityref `json:"campaigns"`


    // Sequences - The list of sequences for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a sequence.
    Sequences []Domainentityref `json:"sequences"`


    // EmailCampaigns - The list of Email campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a Email campaign.
    EmailCampaigns []Domainentityref `json:"emailCampaigns"`


    // SmsCampaigns - The list of SMS campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a SMS campaign.
    SmsCampaigns []Domainentityref `json:"smsCampaigns"`


    // UseTriggeringEntity - If true, the CampaignRuleAction will apply to the same entity that triggered the CampaignRuleCondition.
    UseTriggeringEntity bool `json:"useTriggeringEntity"`

}

// String returns a JSON representation of the model
func (o *Campaignruleactionentities) String() string {
     o.Campaigns = []Domainentityref{{}} 
     o.Sequences = []Domainentityref{{}} 
     o.EmailCampaigns = []Domainentityref{{}} 
     o.SmsCampaigns = []Domainentityref{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleactionentities) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleactionentities

    if CampaignruleactionentitiesMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleactionentitiesMarshalled = true

    return json.Marshal(&struct {
        
        Campaigns []Domainentityref `json:"campaigns"`
        
        Sequences []Domainentityref `json:"sequences"`
        
        EmailCampaigns []Domainentityref `json:"emailCampaigns"`
        
        SmsCampaigns []Domainentityref `json:"smsCampaigns"`
        
        UseTriggeringEntity bool `json:"useTriggeringEntity"`
        *Alias
    }{

        
        Campaigns: []Domainentityref{{}},
        


        
        Sequences: []Domainentityref{{}},
        


        
        EmailCampaigns: []Domainentityref{{}},
        


        
        SmsCampaigns: []Domainentityref{{}},
        


        

        Alias: (*Alias)(u),
    })
}


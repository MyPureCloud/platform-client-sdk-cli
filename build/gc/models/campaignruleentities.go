package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleentitiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleentitiesDud struct { 
    


    

}

// Campaignruleentities
type Campaignruleentities struct { 
    // Campaigns - The list of campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a campaign.
    Campaigns []Domainentityref `json:"campaigns"`


    // Sequences - The list of sequences for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a sequence.
    Sequences []Domainentityref `json:"sequences"`

}

// String returns a JSON representation of the model
func (o *Campaignruleentities) String() string {
     o.Campaigns = []Domainentityref{{}} 
     o.Sequences = []Domainentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleentities) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleentities

    if CampaignruleentitiesMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleentitiesMarshalled = true

    return json.Marshal(&struct {
        
        Campaigns []Domainentityref `json:"campaigns"`
        
        Sequences []Domainentityref `json:"sequences"`
        *Alias
    }{

        
        Campaigns: []Domainentityref{{}},
        


        
        Sequences: []Domainentityref{{}},
        

        Alias: (*Alias)(u),
    })
}


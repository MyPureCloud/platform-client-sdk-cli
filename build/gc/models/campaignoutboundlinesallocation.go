package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignoutboundlinesallocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignoutboundlinesallocationDud struct { 
    


    


    


    

}

// Campaignoutboundlinesallocation
type Campaignoutboundlinesallocation struct { 
    // Campaign - The Campaign
    Campaign Domainentityref `json:"campaign"`


    // CampaignWeight - The relative weight of the campaign
    CampaignWeight int `json:"campaignWeight"`


    // LinesAssigned - The number of lines dynamically assigned to the campaign
    LinesAssigned int `json:"linesAssigned"`


    // LegacyWeight - true if relative weight of the campaign is not explicitly specified, false otherwise
    LegacyWeight bool `json:"legacyWeight"`

}

// String returns a JSON representation of the model
func (o *Campaignoutboundlinesallocation) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignoutboundlinesallocation) MarshalJSON() ([]byte, error) {
    type Alias Campaignoutboundlinesallocation

    if CampaignoutboundlinesallocationMarshalled {
        return []byte("{}"), nil
    }
    CampaignoutboundlinesallocationMarshalled = true

    return json.Marshal(&struct {
        
        Campaign Domainentityref `json:"campaign"`
        
        CampaignWeight int `json:"campaignWeight"`
        
        LinesAssigned int `json:"linesAssigned"`
        
        LegacyWeight bool `json:"legacyWeight"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


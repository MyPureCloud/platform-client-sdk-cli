package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignlinesutilizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignlinesutilizationDud struct { 
    


    

}

// Campaignlinesutilization
type Campaignlinesutilization struct { 
    // AssignedOutboundLines - Number of outbound lines assigned to the campaign
    AssignedOutboundLines int `json:"assignedOutboundLines"`


    // TotalAvailableOutboundLines - Total number of available outbound lines in Campaign's Edge Group or Site
    TotalAvailableOutboundLines int `json:"totalAvailableOutboundLines"`

}

// String returns a JSON representation of the model
func (o *Campaignlinesutilization) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignlinesutilization) MarshalJSON() ([]byte, error) {
    type Alias Campaignlinesutilization

    if CampaignlinesutilizationMarshalled {
        return []byte("{}"), nil
    }
    CampaignlinesutilizationMarshalled = true

    return json.Marshal(&struct {
        
        AssignedOutboundLines int `json:"assignedOutboundLines"`
        
        TotalAvailableOutboundLines int `json:"totalAvailableOutboundLines"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


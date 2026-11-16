package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleoutboundblendingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleoutboundblendingsettingsDud struct { 
    


    

}

// Campaignruleoutboundblendingsettings
type Campaignruleoutboundblendingsettings struct { 
    // Enabled - Whether Enhanced Blending is enabled for the queue.
    Enabled bool `json:"enabled"`


    // CampaignReservationPercentage - The proportion of on-queue agents to reserve for outbound campaign calls. Allowable range 1 - 100 (inclusive).
    CampaignReservationPercentage int `json:"campaignReservationPercentage"`

}

// String returns a JSON representation of the model
func (o *Campaignruleoutboundblendingsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleoutboundblendingsettings) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleoutboundblendingsettings

    if CampaignruleoutboundblendingsettingsMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleoutboundblendingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        CampaignReservationPercentage int `json:"campaignReservationPercentage"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


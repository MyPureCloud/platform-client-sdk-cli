package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutboundblendingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutboundblendingsettingsDud struct { 
    


    

}

// Outboundblendingsettings
type Outboundblendingsettings struct { 
    // Enabled - Whether Enhanced Blending is enabled for the queue.
    Enabled bool `json:"enabled"`


    // CampaignReservationPercentage - The proportion of on-queue agents to reserve for outbound campaign calls. Allowable range 1 - 100 (inclusive).
    CampaignReservationPercentage int `json:"campaignReservationPercentage"`

}

// String returns a JSON representation of the model
func (o *Outboundblendingsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outboundblendingsettings) MarshalJSON() ([]byte, error) {
    type Alias Outboundblendingsettings

    if OutboundblendingsettingsMarshalled {
        return []byte("{}"), nil
    }
    OutboundblendingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        CampaignReservationPercentage int `json:"campaignReservationPercentage"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


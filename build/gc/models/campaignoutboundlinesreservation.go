package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignoutboundlinesreservationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignoutboundlinesreservationDud struct { 
    


    

}

// Campaignoutboundlinesreservation
type Campaignoutboundlinesreservation struct { 
    // Campaign - The Campaign
    Campaign Domainentityref `json:"campaign"`


    // LinesReserved - The number of lines reserved for the campaign
    LinesReserved int `json:"linesReserved"`

}

// String returns a JSON representation of the model
func (o *Campaignoutboundlinesreservation) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignoutboundlinesreservation) MarshalJSON() ([]byte, error) {
    type Alias Campaignoutboundlinesreservation

    if CampaignoutboundlinesreservationMarshalled {
        return []byte("{}"), nil
    }
    CampaignoutboundlinesreservationMarshalled = true

    return json.Marshal(&struct {
        
        Campaign Domainentityref `json:"campaign"`
        
        LinesReserved int `json:"linesReserved"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignstatsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignstatsDud struct { 
    ContactRate Connectrate `json:"contactRate"`


    IdleAgents int `json:"idleAgents"`


    EffectiveIdleAgents float64 `json:"effectiveIdleAgents"`


    AdjustedCallsPerAgent float64 `json:"adjustedCallsPerAgent"`


    OutstandingCalls int `json:"outstandingCalls"`


    ScheduledCalls int `json:"scheduledCalls"`


    TimeZoneRescheduledCalls int `json:"timeZoneRescheduledCalls"`


    FilteredOutContactsCount int `json:"filteredOutContactsCount"`


    LinesUtilization Campaignlinesutilization `json:"linesUtilization"`

}

// Campaignstats
type Campaignstats struct { 
    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Campaignstats) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignstats) MarshalJSON() ([]byte, error) {
    type Alias Campaignstats

    if CampaignstatsMarshalled {
        return []byte("{}"), nil
    }
    CampaignstatsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigndiagnosticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigndiagnosticsDud struct { 
    CallableContacts Callablecontactsdiagnostic `json:"callableContacts"`


    QueueUtilizationDiagnostic Queueutilizationdiagnostic `json:"queueUtilizationDiagnostic"`


    RuleSetDiagnostics []Rulesetdiagnostic `json:"ruleSetDiagnostics"`


    OutstandingInteractionsCount int `json:"outstandingInteractionsCount"`


    ScheduledInteractionsCount int `json:"scheduledInteractionsCount"`


    TimeZoneRescheduledCallsCount int `json:"timeZoneRescheduledCallsCount"`

}

// Campaigndiagnostics
type Campaigndiagnostics struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Campaigndiagnostics) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigndiagnostics) MarshalJSON() ([]byte, error) {
    type Alias Campaigndiagnostics

    if CampaigndiagnosticsMarshalled {
        return []byte("{}"), nil
    }
    CampaigndiagnosticsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


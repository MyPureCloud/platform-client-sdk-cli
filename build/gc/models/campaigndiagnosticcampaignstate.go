package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigndiagnosticcampaignstateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigndiagnosticcampaignstateDud struct { 
    


    


    


    

}

// Campaigndiagnosticcampaignstate
type Campaigndiagnosticcampaignstate struct { 
    // State - Campaign status
    State string `json:"state"`


    // DateStart - Start datetime of the state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End datetime of the state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // DialingMode - Dialing mode for the campaign state (e.g., \"power\", \"preview\", \"predictive\")
    DialingMode string `json:"dialingMode"`

}

// String returns a JSON representation of the model
func (o *Campaigndiagnosticcampaignstate) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigndiagnosticcampaignstate) MarshalJSON() ([]byte, error) {
    type Alias Campaigndiagnosticcampaignstate

    if CampaigndiagnosticcampaignstateMarshalled {
        return []byte("{}"), nil
    }
    CampaigndiagnosticcampaignstateMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        DialingMode string `json:"dialingMode"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


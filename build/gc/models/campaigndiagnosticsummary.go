package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigndiagnosticsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigndiagnosticsummaryDud struct { 
    


    


    


    


    


    


    

}

// Campaigndiagnosticsummary
type Campaigndiagnosticsummary struct { 
    // CampaignId - Campaign ID
    CampaignId string `json:"campaignId"`


    // DateStart - Start of the interval. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End of the interval. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // CampaignStates - Array of campaign states
    CampaignStates []Campaigndiagnosticcampaignstate `json:"campaignStates"`


    // CampaignInfo - Array of diagnostic windows
    CampaignInfo []Campaigndiagnosticwindow `json:"campaignInfo"`


    // CampaignHealthStates - Array of campaign health states
    CampaignHealthStates []Campaigndiagnosticcampaignhealthstate `json:"campaignHealthStates"`


    // ConfigChanges - Configuration changes occurring within the time window
    ConfigChanges []Campaigndiagnosticconfigchange `json:"configChanges"`

}

// String returns a JSON representation of the model
func (o *Campaigndiagnosticsummary) String() string {
    
    
    
     o.CampaignStates = []Campaigndiagnosticcampaignstate{{}} 
     o.CampaignInfo = []Campaigndiagnosticwindow{{}} 
     o.CampaignHealthStates = []Campaigndiagnosticcampaignhealthstate{{}} 
     o.ConfigChanges = []Campaigndiagnosticconfigchange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigndiagnosticsummary) MarshalJSON() ([]byte, error) {
    type Alias Campaigndiagnosticsummary

    if CampaigndiagnosticsummaryMarshalled {
        return []byte("{}"), nil
    }
    CampaigndiagnosticsummaryMarshalled = true

    return json.Marshal(&struct {
        
        CampaignId string `json:"campaignId"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        CampaignStates []Campaigndiagnosticcampaignstate `json:"campaignStates"`
        
        CampaignInfo []Campaigndiagnosticwindow `json:"campaignInfo"`
        
        CampaignHealthStates []Campaigndiagnosticcampaignhealthstate `json:"campaignHealthStates"`
        
        ConfigChanges []Campaigndiagnosticconfigchange `json:"configChanges"`
        *Alias
    }{

        


        


        


        
        CampaignStates: []Campaigndiagnosticcampaignstate{{}},
        


        
        CampaignInfo: []Campaigndiagnosticwindow{{}},
        


        
        CampaignHealthStates: []Campaigndiagnosticcampaignhealthstate{{}},
        


        
        ConfigChanges: []Campaigndiagnosticconfigchange{{}},
        

        Alias: (*Alias)(u),
    })
}


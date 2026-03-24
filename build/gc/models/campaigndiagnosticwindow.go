package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigndiagnosticwindowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigndiagnosticwindowDud struct { 
    


    


    

}

// Campaigndiagnosticwindow
type Campaigndiagnosticwindow struct { 
    // State - Name of informational attribute
    State string `json:"state"`


    // DateStart - Start datetime of the window. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End datetime of the window. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`

}

// String returns a JSON representation of the model
func (o *Campaigndiagnosticwindow) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigndiagnosticwindow) MarshalJSON() ([]byte, error) {
    type Alias Campaigndiagnosticwindow

    if CampaigndiagnosticwindowMarshalled {
        return []byte("{}"), nil
    }
    CampaigndiagnosticwindowMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


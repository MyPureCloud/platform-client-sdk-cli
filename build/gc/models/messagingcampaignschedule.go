package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingcampaignscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingcampaignscheduleDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Messagingcampaignschedule
type Messagingcampaignschedule struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Intervals - A list of intervals during which to run the associated Campaign.
    Intervals []Scheduleinterval `json:"intervals"`


    // TimeZone - The time zone for this messaging campaign schedule.
    TimeZone string `json:"timeZone"`


    // MessagingCampaign - The Campaign that this messaging campaign schedule is for.
    MessagingCampaign Divisioneddomainentityref `json:"messagingCampaign"`


    

}

// String returns a JSON representation of the model
func (o *Messagingcampaignschedule) String() string {
    
    
     o.Intervals = []Scheduleinterval{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingcampaignschedule) MarshalJSON() ([]byte, error) {
    type Alias Messagingcampaignschedule

    if MessagingcampaignscheduleMarshalled {
        return []byte("{}"), nil
    }
    MessagingcampaignscheduleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Intervals []Scheduleinterval `json:"intervals"`
        
        TimeZone string `json:"timeZone"`
        
        MessagingCampaign Divisioneddomainentityref `json:"messagingCampaign"`
        *Alias
    }{

        


        


        


        


        


        
        Intervals: []Scheduleinterval{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


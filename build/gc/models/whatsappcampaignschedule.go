package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappcampaignscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappcampaignscheduleDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Whatsappcampaignschedule
type Whatsappcampaignschedule struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Intervals - A list of intervals during which to run the associated Campaign.
    Intervals []Scheduleinterval `json:"intervals"`


    // Recurrences - Recurring schedules of the campaign
    Recurrences []Reoccurrence `json:"recurrences"`


    // TimeZone - The time zone for this WhatsApp campaign schedule. Defaults to UTC if empty or not provided. See here for a list of valid time zones https://www.iana.org/time-zones
    TimeZone string `json:"timeZone"`


    // WhatsAppCampaign - The Campaign that this WhatsApp campaign schedule is for.
    WhatsAppCampaign Divisioneddomainentityref `json:"whatsAppCampaign"`


    

}

// String returns a JSON representation of the model
func (o *Whatsappcampaignschedule) String() string {
    
    
     o.Intervals = []Scheduleinterval{{}} 
     o.Recurrences = []Reoccurrence{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappcampaignschedule) MarshalJSON() ([]byte, error) {
    type Alias Whatsappcampaignschedule

    if WhatsappcampaignscheduleMarshalled {
        return []byte("{}"), nil
    }
    WhatsappcampaignscheduleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Intervals []Scheduleinterval `json:"intervals"`
        
        Recurrences []Reoccurrence `json:"recurrences"`
        
        TimeZone string `json:"timeZone"`
        
        WhatsAppCampaign Divisioneddomainentityref `json:"whatsAppCampaign"`
        *Alias
    }{

        


        


        


        


        


        
        Intervals: []Scheduleinterval{{}},
        


        
        Recurrences: []Reoccurrence{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


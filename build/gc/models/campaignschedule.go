package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignscheduleDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Campaignschedule
type Campaignschedule struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Intervals - A list of intervals during which to run the associated Campaign.
    Intervals []Scheduleinterval `json:"intervals"`


    // Recurrences - Recurring schedules of the campaign
    Recurrences []Reoccurrence `json:"recurrences"`


    // TimeZone - The time zone for this CampaignSchedule. Defaults to UTC if empty or not provided. See here for a list of valid time zones https://www.iana.org/time-zones
    TimeZone string `json:"timeZone"`


    // Campaign - The Campaign that this CampaignSchedule is for.
    Campaign Divisioneddomainentityref `json:"campaign"`


    

}

// String returns a JSON representation of the model
func (o *Campaignschedule) String() string {
    
    
     o.Intervals = []Scheduleinterval{{}} 
     o.Recurrences = []Reoccurrence{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignschedule) MarshalJSON() ([]byte, error) {
    type Alias Campaignschedule

    if CampaignscheduleMarshalled {
        return []byte("{}"), nil
    }
    CampaignscheduleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Intervals []Scheduleinterval `json:"intervals"`
        
        Recurrences []Reoccurrence `json:"recurrences"`
        
        TimeZone string `json:"timeZone"`
        
        Campaign Divisioneddomainentityref `json:"campaign"`
        *Alias
    }{

        


        


        


        


        


        
        Intervals: []Scheduleinterval{{}},
        


        
        Recurrences: []Reoccurrence{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


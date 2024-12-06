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


    // TimeZone - The time zone for this CampaignSchedule. For example, Africa/Abidjan.
    TimeZone string `json:"timeZone"`


    // Campaign - The Campaign that this CampaignSchedule is for.
    Campaign Divisioneddomainentityref `json:"campaign"`


    

}

// String returns a JSON representation of the model
func (o *Campaignschedule) String() string {
    
    
     o.Intervals = []Scheduleinterval{{}} 
    
    

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
        
        TimeZone string `json:"timeZone"`
        
        Campaign Divisioneddomainentityref `json:"campaign"`
        *Alias
    }{

        


        


        


        


        


        
        Intervals: []Scheduleinterval{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


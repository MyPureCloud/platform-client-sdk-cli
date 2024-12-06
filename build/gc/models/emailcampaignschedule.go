package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailcampaignscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailcampaignscheduleDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Emailcampaignschedule
type Emailcampaignschedule struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Intervals - A list of intervals during which to run the associated Campaign.
    Intervals []Scheduleinterval `json:"intervals"`


    // TimeZone - The time zone for this email campaign schedule.
    TimeZone string `json:"timeZone"`


    // EmailCampaign - The Campaign that this email campaign schedule is for.
    EmailCampaign Divisioneddomainentityref `json:"emailCampaign"`


    

}

// String returns a JSON representation of the model
func (o *Emailcampaignschedule) String() string {
    
    
     o.Intervals = []Scheduleinterval{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailcampaignschedule) MarshalJSON() ([]byte, error) {
    type Alias Emailcampaignschedule

    if EmailcampaignscheduleMarshalled {
        return []byte("{}"), nil
    }
    EmailcampaignscheduleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Intervals []Scheduleinterval `json:"intervals"`
        
        TimeZone string `json:"timeZone"`
        
        EmailCampaign Divisioneddomainentityref `json:"emailCampaign"`
        *Alias
    }{

        


        


        


        


        


        
        Intervals: []Scheduleinterval{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


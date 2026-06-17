package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidscheduledactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidscheduledactivityDud struct { 
    


    


    


    


    


    

}

// Schedulebidscheduledactivity
type Schedulebidscheduledactivity struct { 
    // StartDate - The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of this activity in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Description - The description of this activity
    Description string `json:"description"`


    // ActivityCategory - The activity code's category
    ActivityCategory string `json:"activityCategory"`


    // ActivityCodeId - The ID of the activity code associated with this activity
    ActivityCodeId string `json:"activityCodeId"`


    // Paid - Whether this activity is paid
    Paid bool `json:"paid"`

}

// String returns a JSON representation of the model
func (o *Schedulebidscheduledactivity) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidscheduledactivity) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidscheduledactivity

    if SchedulebidscheduledactivityMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidscheduledactivityMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Description string `json:"description"`
        
        ActivityCategory string `json:"activityCategory"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


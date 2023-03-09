package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningslotscheduleactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningslotscheduleactivityDud struct { 
    


    


    


    


    


    


    


    

}

// Learningslotscheduleactivity
type Learningslotscheduleactivity struct { 
    // DateStart - The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // LengthMinutes - The length of this activity in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Description - The description of this activity
    Description string `json:"description"`


    // ActivityCodeId - The ID of the activity code associated with this activity
    ActivityCodeId string `json:"activityCodeId"`


    // Paid - Whether this activity is paid
    Paid bool `json:"paid"`


    // TimeOffRequestId - The ID of the time off request associated with this activity, if applicable
    TimeOffRequestId string `json:"timeOffRequestId"`


    // ExternalActivityId - The ID of the external activity associated with this activity, if applicable
    ExternalActivityId string `json:"externalActivityId"`


    // ExternalActivityType - The type of the external activity associated with this activity, if applicable
    ExternalActivityType string `json:"externalActivityType"`

}

// String returns a JSON representation of the model
func (o *Learningslotscheduleactivity) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningslotscheduleactivity) MarshalJSON() ([]byte, error) {
    type Alias Learningslotscheduleactivity

    if LearningslotscheduleactivityMarshalled {
        return []byte("{}"), nil
    }
    LearningslotscheduleactivityMarshalled = true

    return json.Marshal(&struct {
        
        DateStart time.Time `json:"dateStart"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Description string `json:"description"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        
        TimeOffRequestId string `json:"timeOffRequestId"`
        
        ExternalActivityId string `json:"externalActivityId"`
        
        ExternalActivityType string `json:"externalActivityType"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


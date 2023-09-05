package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentscheduleactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentscheduleactivityDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Buagentscheduleactivity
type Buagentscheduleactivity struct { 
    // StartDate - The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of this activity in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Description - The description of this activity
    Description string `json:"description"`


    // ActivityCodeId - The ID of the activity code associated with this activity
    ActivityCodeId string `json:"activityCodeId"`


    // Paid - Whether this activity is paid
    Paid bool `json:"paid"`


    // PayableMinutes - Payable minutes for this activity
    PayableMinutes int `json:"payableMinutes"`


    // TimeOffRequestId - The ID of the time off request associated with this activity, if applicable
    TimeOffRequestId string `json:"timeOffRequestId"`


    // TimeOffRequestSyncVersion - The sync version of the partial day time off request for which the scheduled activity is associated, if applicable
    TimeOffRequestSyncVersion int `json:"timeOffRequestSyncVersion"`


    // ExternalActivityId - The ID of the external activity associated with this activity, if applicable
    ExternalActivityId string `json:"externalActivityId"`


    // ExternalActivityType - The type of the external activity associated with this activity, if applicable
    ExternalActivityType string `json:"externalActivityType"`

}

// String returns a JSON representation of the model
func (o *Buagentscheduleactivity) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentscheduleactivity) MarshalJSON() ([]byte, error) {
    type Alias Buagentscheduleactivity

    if BuagentscheduleactivityMarshalled {
        return []byte("{}"), nil
    }
    BuagentscheduleactivityMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Description string `json:"description"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        
        PayableMinutes int `json:"payableMinutes"`
        
        TimeOffRequestId string `json:"timeOffRequestId"`
        
        TimeOffRequestSyncVersion int `json:"timeOffRequestSyncVersion"`
        
        ExternalActivityId string `json:"externalActivityId"`
        
        ExternalActivityType string `json:"externalActivityType"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


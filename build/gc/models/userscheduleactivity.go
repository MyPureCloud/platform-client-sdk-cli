package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserscheduleactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserscheduleactivityDud struct { 
    


    


    


    


    


    


    

}

// Userscheduleactivity
type Userscheduleactivity struct { 
    // ActivityCodeId - The id for the activity code.  Look up a map of activity codes with the activities route
    ActivityCodeId string `json:"activityCodeId"`


    // StartDate - Start time in UTC for this activity, in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthInMinutes - Length in minutes for this activity
    LengthInMinutes int `json:"lengthInMinutes"`


    // Description - Description for this activity
    Description string `json:"description"`


    // CountsAsPaidTime - Whether this activity is paid
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // IsDstFallback - Whether this activity spans a DST fallback
    IsDstFallback bool `json:"isDstFallback"`


    // TimeOffRequestId - Time off request id of this activity
    TimeOffRequestId string `json:"timeOffRequestId"`

}

// String returns a JSON representation of the model
func (o *Userscheduleactivity) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userscheduleactivity) MarshalJSON() ([]byte, error) {
    type Alias Userscheduleactivity

    if UserscheduleactivityMarshalled {
        return []byte("{}"), nil
    }
    UserscheduleactivityMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCodeId string `json:"activityCodeId"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        Description string `json:"description"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        IsDstFallback bool `json:"isDstFallback"`
        
        TimeOffRequestId string `json:"timeOffRequestId"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


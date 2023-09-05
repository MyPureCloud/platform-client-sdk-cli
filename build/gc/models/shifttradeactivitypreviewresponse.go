package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradeactivitypreviewresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradeactivitypreviewresponseDud struct { 
    


    


    


    


    

}

// Shifttradeactivitypreviewresponse
type Shifttradeactivitypreviewresponse struct { 
    // StartDate - The start date and time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length in minutes of this activity
    LengthMinutes int `json:"lengthMinutes"`


    // ActivityCodeId - The ID of the activity code for this activity
    ActivityCodeId string `json:"activityCodeId"`


    // CountsAsPaidTime - Whether this activity counts as paid time
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // PayableMinutes - Payable minutes for this activity
    PayableMinutes int `json:"payableMinutes"`

}

// String returns a JSON representation of the model
func (o *Shifttradeactivitypreviewresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradeactivitypreviewresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradeactivitypreviewresponse

    if ShifttradeactivitypreviewresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradeactivitypreviewresponseMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        PayableMinutes int `json:"payableMinutes"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


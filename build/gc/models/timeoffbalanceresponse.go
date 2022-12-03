package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffbalanceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffbalanceresponseDud struct { 
    


    


    


    


    

}

// Timeoffbalanceresponse
type Timeoffbalanceresponse struct { 
    // ActivityCodeId - The ID for activity code associated with time off balance
    ActivityCodeId string `json:"activityCodeId"`


    // HrisTimeOffTypeId - The ID of the time off type configured in HRIS integration
    HrisTimeOffTypeId string `json:"hrisTimeOffTypeId"`


    // HrisTimeOffTypeSecondaryId - The secondary ID of the time off type configured in HRIS integration
    HrisTimeOffTypeSecondaryId string `json:"hrisTimeOffTypeSecondaryId"`


    // StartDate - The Start date of the requested date range. The end date is determined by the size of interval list. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // BalanceMinutesPerDay - The list of available time off balance values in minutes for each day
    BalanceMinutesPerDay []int `json:"balanceMinutesPerDay"`

}

// String returns a JSON representation of the model
func (o *Timeoffbalanceresponse) String() string {
    
    
    
    
     o.BalanceMinutesPerDay = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffbalanceresponse) MarshalJSON() ([]byte, error) {
    type Alias Timeoffbalanceresponse

    if TimeoffbalanceresponseMarshalled {
        return []byte("{}"), nil
    }
    TimeoffbalanceresponseMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCodeId string `json:"activityCodeId"`
        
        HrisTimeOffTypeId string `json:"hrisTimeOffTypeId"`
        
        HrisTimeOffTypeSecondaryId string `json:"hrisTimeOffTypeSecondaryId"`
        
        StartDate time.Time `json:"startDate"`
        
        BalanceMinutesPerDay []int `json:"balanceMinutesPerDay"`
        *Alias
    }{

        


        


        


        


        
        BalanceMinutesPerDay: []int{0},
        

        Alias: (*Alias)(u),
    })
}


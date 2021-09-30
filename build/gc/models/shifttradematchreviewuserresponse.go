package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradematchreviewuserresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradematchreviewuserresponseDud struct { 
    


    


    


    


    

}

// Shifttradematchreviewuserresponse
type Shifttradematchreviewuserresponse struct { 
    // WeeklyMinimumPaidMinutes - The minimum weekly paid minutes for this user per the work plan tied to the agent schedule
    WeeklyMinimumPaidMinutes int `json:"weeklyMinimumPaidMinutes"`


    // WeeklyMaximumPaidMinutes - The maximum weekly paid minutes for this user per the work plan tied to the agent schedule
    WeeklyMaximumPaidMinutes int `json:"weeklyMaximumPaidMinutes"`


    // PreTradeSchedulePaidMinutes - The paid minutes on the week schedule for this user prior to the shift trade
    PreTradeSchedulePaidMinutes int `json:"preTradeSchedulePaidMinutes"`


    // PostTradeSchedulePaidMinutes - The paid minutes on the week schedule for this user if the shift trade is approved
    PostTradeSchedulePaidMinutes int `json:"postTradeSchedulePaidMinutes"`


    // PostTradeNewShift - Preview of what the shift will look like for the opposite side of this trade after the match is approved
    PostTradeNewShift Shifttradepreviewresponse `json:"postTradeNewShift"`

}

// String returns a JSON representation of the model
func (o *Shifttradematchreviewuserresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradematchreviewuserresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradematchreviewuserresponse

    if ShifttradematchreviewuserresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradematchreviewuserresponseMarshalled = true

    return json.Marshal(&struct { 
        WeeklyMinimumPaidMinutes int `json:"weeklyMinimumPaidMinutes"`
        
        WeeklyMaximumPaidMinutes int `json:"weeklyMaximumPaidMinutes"`
        
        PreTradeSchedulePaidMinutes int `json:"preTradeSchedulePaidMinutes"`
        
        PostTradeSchedulePaidMinutes int `json:"postTradeSchedulePaidMinutes"`
        
        PostTradeNewShift Shifttradepreviewresponse `json:"postTradeNewShift"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradenotificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradenotificationDud struct { 
    


    


    


    


    


    


    


    

}

// Shifttradenotification
type Shifttradenotification struct { 
    // WeekDate - The start week date of the initiating shift of the shift trade in yyyy-MM-dd format
    WeekDate string `json:"weekDate"`


    // TradeId - The ID of the shift trade
    TradeId string `json:"tradeId"`


    // OneSided - Whether this is a one sided shift trade
    OneSided bool `json:"oneSided"`


    // NewState - The new state of the shift trade, null if there was no change
    NewState string `json:"newState"`


    // InitiatingUser - The user who initiated the shift trade
    InitiatingUser Userreference `json:"initiatingUser"`


    // InitiatingShiftDate - The start date and time of the initiating shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    InitiatingShiftDate time.Time `json:"initiatingShiftDate"`


    // ReceivingUser - The user on the receiving side of this shift trade (null if not matched)
    ReceivingUser Userreference `json:"receivingUser"`


    // ReceivingShiftDate - The start date and time of the receiving shift (null if not matched or if one-sided. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReceivingShiftDate time.Time `json:"receivingShiftDate"`

}

// String returns a JSON representation of the model
func (o *Shifttradenotification) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradenotification) MarshalJSON() ([]byte, error) {
    type Alias Shifttradenotification

    if ShifttradenotificationMarshalled {
        return []byte("{}"), nil
    }
    ShifttradenotificationMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate string `json:"weekDate"`
        
        TradeId string `json:"tradeId"`
        
        OneSided bool `json:"oneSided"`
        
        NewState string `json:"newState"`
        
        InitiatingUser Userreference `json:"initiatingUser"`
        
        InitiatingShiftDate time.Time `json:"initiatingShiftDate"`
        
        ReceivingUser Userreference `json:"receivingUser"`
        
        ReceivingShiftDate time.Time `json:"receivingShiftDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftnotificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftnotificationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Alternativeshiftnotification
type Alternativeshiftnotification struct { 
    


    // WeekDate - The start date of the schedule with which this trade is associated. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // Granularity - The granularity of alternative shifts to be traded
    Granularity string `json:"granularity"`


    // NewState - The new state of the alternative shift trade, null if there was no change
    NewState string `json:"newState"`


    // InitiatingUser - The user who initiated the alternative shift trade
    InitiatingUser Userreference `json:"initiatingUser"`


    // InitiatingShiftDate - The start date and time of the initiating shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    InitiatingShiftDate time.Time `json:"initiatingShiftDate"`


    // ReceivingUser - The user on the receiving side of this alternative shift trade
    ReceivingUser Userreference `json:"receivingUser"`


    // ReceivingShiftDate - The start date and time of the receiving alternative shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReceivingShiftDate time.Time `json:"receivingShiftDate"`


    

}

// String returns a JSON representation of the model
func (o *Alternativeshiftnotification) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftnotification) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftnotification

    if AlternativeshiftnotificationMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftnotificationMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        Granularity string `json:"granularity"`
        
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


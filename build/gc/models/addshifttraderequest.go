package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddshifttraderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddshifttraderequestDud struct { 
    


    


    


    


    

}

// Addshifttraderequest
type Addshifttraderequest struct { 
    // ScheduleId - The ID of the schedule to which the initiating and receiving shifts belong
    ScheduleId string `json:"scheduleId"`


    // InitiatingShiftId - The ID of the shift that the initiating user wants to give up
    InitiatingShiftId string `json:"initiatingShiftId"`


    // ReceivingUserId - The ID of the user to whom to send the request (for use in direct trade requests)
    ReceivingUserId string `json:"receivingUserId"`


    // Expiration - When this shift trade request should expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Expiration time.Time `json:"expiration"`


    // AcceptableIntervals
    AcceptableIntervals []string `json:"acceptableIntervals"`

}

// String returns a JSON representation of the model
func (o *Addshifttraderequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AcceptableIntervals = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addshifttraderequest) MarshalJSON() ([]byte, error) {
    type Alias Addshifttraderequest

    if AddshifttraderequestMarshalled {
        return []byte("{}"), nil
    }
    AddshifttraderequestMarshalled = true

    return json.Marshal(&struct { 
        ScheduleId string `json:"scheduleId"`
        
        InitiatingShiftId string `json:"initiatingShiftId"`
        
        ReceivingUserId string `json:"receivingUserId"`
        
        Expiration time.Time `json:"expiration"`
        
        AcceptableIntervals []string `json:"acceptableIntervals"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        AcceptableIntervals: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}


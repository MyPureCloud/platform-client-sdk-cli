package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttraderesponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttraderesponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Shifttraderesponse
type Shifttraderesponse struct { 
    // Id - The ID of this shift trade
    Id string `json:"id"`


    // WeekDate - The start week date of the associated schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // Schedule - A reference to the associated schedule
    Schedule Buschedulereferenceformuroute `json:"schedule"`


    // State - The state of this shift trade
    State string `json:"state"`


    // InitiatingUser - The user who initiated this trade
    InitiatingUser Userreference `json:"initiatingUser"`


    // InitiatingShiftId - The ID of the shift offered for trade by the initiating user
    InitiatingShiftId string `json:"initiatingShiftId"`


    // InitiatingShiftStart - The start date/time of the shift being offered for trade. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    InitiatingShiftStart time.Time `json:"initiatingShiftStart"`


    // InitiatingShiftEnd - The end date/time of the shift being offered for trade. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    InitiatingShiftEnd time.Time `json:"initiatingShiftEnd"`


    // ReceivingUser - The user matching the trade, or if the state is not Matched, the user to whom the trade request was sent
    ReceivingUser Userreference `json:"receivingUser"`


    // ReceivingShiftId - The ID of the shift being exchanged for the initiating shift, null if the receiving user is picking up a shift
    ReceivingShiftId string `json:"receivingShiftId"`


    // ReceivingShiftStart - The start date/time of the receiving shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReceivingShiftStart time.Time `json:"receivingShiftStart"`


    // ReceivingShiftEnd - The end date/time of the receiving shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReceivingShiftEnd time.Time `json:"receivingShiftEnd"`


    // Expiration - When this shift trade offer will expire if not matched or approved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Expiration time.Time `json:"expiration"`


    // OneSided - Whether this is a one-sided shift trade (e.g. the initiating user is not asking for a shift in return)
    OneSided bool `json:"oneSided"`


    // AcceptableIntervals - Time frames when the initiating user is willing to accept trades.  Empty means giving up the shift. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    AcceptableIntervals []string `json:"acceptableIntervals"`


    // ReviewedBy - The user who reviewed this shift trade
    ReviewedBy Userreference `json:"reviewedBy"`


    // ReviewedDate - The timestamp when this shift trade was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReviewedDate time.Time `json:"reviewedDate"`


    // Metadata - Version data for this trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Shifttraderesponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AcceptableIntervals = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttraderesponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttraderesponse

    if ShifttraderesponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttraderesponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        
        Schedule Buschedulereferenceformuroute `json:"schedule"`
        
        State string `json:"state"`
        
        InitiatingUser Userreference `json:"initiatingUser"`
        
        InitiatingShiftId string `json:"initiatingShiftId"`
        
        InitiatingShiftStart time.Time `json:"initiatingShiftStart"`
        
        InitiatingShiftEnd time.Time `json:"initiatingShiftEnd"`
        
        ReceivingUser Userreference `json:"receivingUser"`
        
        ReceivingShiftId string `json:"receivingShiftId"`
        
        ReceivingShiftStart time.Time `json:"receivingShiftStart"`
        
        ReceivingShiftEnd time.Time `json:"receivingShiftEnd"`
        
        Expiration time.Time `json:"expiration"`
        
        OneSided bool `json:"oneSided"`
        
        AcceptableIntervals []string `json:"acceptableIntervals"`
        
        ReviewedBy Userreference `json:"reviewedBy"`
        
        ReviewedDate time.Time `json:"reviewedDate"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        AcceptableIntervals: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


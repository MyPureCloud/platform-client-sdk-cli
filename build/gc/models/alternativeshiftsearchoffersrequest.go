package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftsearchoffersrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftsearchoffersrequestDud struct { 
    


    


    


    

}

// Alternativeshiftsearchoffersrequest
type Alternativeshiftsearchoffersrequest struct { 
    // Schedule - The existing schedule being used to find alternative shift offers
    Schedule Alternativeshiftschedulelookup `json:"schedule"`


    // QueryWeekDate - The start date for the week in this schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    QueryWeekDate time.Time `json:"queryWeekDate"`


    // InitiatingShift - The shift a user puts up for alternative shift offers
    InitiatingShift Initiatingalternativeshift `json:"initiatingShift"`


    // AcceptableIntervals - The acceptable intervals in offers. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    AcceptableIntervals []string `json:"acceptableIntervals"`

}

// String returns a JSON representation of the model
func (o *Alternativeshiftsearchoffersrequest) String() string {
    
    
    
     o.AcceptableIntervals = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftsearchoffersrequest) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftsearchoffersrequest

    if AlternativeshiftsearchoffersrequestMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftsearchoffersrequestMarshalled = true

    return json.Marshal(&struct {
        
        Schedule Alternativeshiftschedulelookup `json:"schedule"`
        
        QueryWeekDate time.Time `json:"queryWeekDate"`
        
        InitiatingShift Initiatingalternativeshift `json:"initiatingShift"`
        
        AcceptableIntervals []string `json:"acceptableIntervals"`
        *Alias
    }{

        


        


        


        
        AcceptableIntervals: []string{""},
        

        Alias: (*Alias)(u),
    })
}


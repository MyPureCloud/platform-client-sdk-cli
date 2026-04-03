package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddshifttradejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddshifttradejobrequestDud struct { 
    


    


    


    

}

// Addshifttradejobrequest
type Addshifttradejobrequest struct { 
    // InitiatingShift - The shift that the initiating user wants to give up in this trade
    InitiatingShift Initiatingshiftrequestitem `json:"initiatingShift"`


    // AcceptableIntervals - Time frames when the initiating user is willing to accept a shift in exchange. Empty means giving up the shift without taking on another one
    AcceptableIntervals []Requireddaterange `json:"acceptableIntervals"`


    // Target - Optional shift trade target, can be used for example for direct user to user trade
    Target Shifttradetargetrequestitem `json:"target"`


    // ExpirationDate - When this shift trade will expire. Date time is represented as an ISO-8601 string
    ExpirationDate time.Time `json:"expirationDate"`

}

// String returns a JSON representation of the model
func (o *Addshifttradejobrequest) String() string {
    
     o.AcceptableIntervals = []Requireddaterange{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addshifttradejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Addshifttradejobrequest

    if AddshifttradejobrequestMarshalled {
        return []byte("{}"), nil
    }
    AddshifttradejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        InitiatingShift Initiatingshiftrequestitem `json:"initiatingShift"`
        
        AcceptableIntervals []Requireddaterange `json:"acceptableIntervals"`
        
        Target Shifttradetargetrequestitem `json:"target"`
        
        ExpirationDate time.Time `json:"expirationDate"`
        *Alias
    }{

        


        
        AcceptableIntervals: []Requireddaterange{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}


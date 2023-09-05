package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestlookupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestlookupDud struct { 
    


    

}

// Timeoffrequestlookup
type Timeoffrequestlookup struct { 
    // TimeOffRequestId - The ID of the time off request
    TimeOffRequestId string `json:"timeOffRequestId"`


    // UserId - The ID of the user to whom the time off request belongs
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestlookup) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestlookup) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestlookup

    if TimeoffrequestlookupMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestlookupMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffRequestId string `json:"timeOffRequestId"`
        
        UserId string `json:"userId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


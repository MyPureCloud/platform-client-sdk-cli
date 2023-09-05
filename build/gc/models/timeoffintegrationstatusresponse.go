package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffintegrationstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffintegrationstatusresponseDud struct { 
    


    

}

// Timeoffintegrationstatusresponse
type Timeoffintegrationstatusresponse struct { 
    // TimeOffRequest - The time off request associated with this integration status
    TimeOffRequest Timeoffrequestreference `json:"timeOffRequest"`


    // IntegrationStatus - The value of integration status for the time off request
    IntegrationStatus string `json:"integrationStatus"`

}

// String returns a JSON representation of the model
func (o *Timeoffintegrationstatusresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffintegrationstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Timeoffintegrationstatusresponse

    if TimeoffintegrationstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    TimeoffintegrationstatusresponseMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffRequest Timeoffrequestreference `json:"timeOffRequest"`
        
        IntegrationStatus string `json:"integrationStatus"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


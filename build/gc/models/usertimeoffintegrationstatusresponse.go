package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsertimeoffintegrationstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsertimeoffintegrationstatusresponseDud struct { 
    


    


    

}

// Usertimeoffintegrationstatusresponse
type Usertimeoffintegrationstatusresponse struct { 
    // TimeOffRequest - The time off request associated with this integration status
    TimeOffRequest Timeoffrequestreference `json:"timeOffRequest"`


    // IntegrationStatus - The value of integration status for the time off request
    IntegrationStatus string `json:"integrationStatus"`


    // User - The user to whom the time off request belongs
    User Userreference `json:"user"`

}

// String returns a JSON representation of the model
func (o *Usertimeoffintegrationstatusresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usertimeoffintegrationstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Usertimeoffintegrationstatusresponse

    if UsertimeoffintegrationstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    UsertimeoffintegrationstatusresponseMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffRequest Timeoffrequestreference `json:"timeOffRequest"`
        
        IntegrationStatus string `json:"integrationStatus"`
        
        User Userreference `json:"user"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


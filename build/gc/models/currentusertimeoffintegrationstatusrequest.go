package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CurrentusertimeoffintegrationstatusrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CurrentusertimeoffintegrationstatusrequestDud struct { 
    

}

// Currentusertimeoffintegrationstatusrequest
type Currentusertimeoffintegrationstatusrequest struct { 
    // TimeOffRequestIds - A list of time off request IDs
    TimeOffRequestIds []string `json:"timeOffRequestIds"`

}

// String returns a JSON representation of the model
func (o *Currentusertimeoffintegrationstatusrequest) String() string {
     o.TimeOffRequestIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Currentusertimeoffintegrationstatusrequest) MarshalJSON() ([]byte, error) {
    type Alias Currentusertimeoffintegrationstatusrequest

    if CurrentusertimeoffintegrationstatusrequestMarshalled {
        return []byte("{}"), nil
    }
    CurrentusertimeoffintegrationstatusrequestMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffRequestIds []string `json:"timeOffRequestIds"`
        *Alias
    }{

        
        TimeOffRequestIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}


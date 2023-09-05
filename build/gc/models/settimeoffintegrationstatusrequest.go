package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SettimeoffintegrationstatusrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SettimeoffintegrationstatusrequestDud struct { 
    

}

// Settimeoffintegrationstatusrequest
type Settimeoffintegrationstatusrequest struct { 
    // IntegrationStatus - The integration status value for the time off request
    IntegrationStatus string `json:"integrationStatus"`

}

// String returns a JSON representation of the model
func (o *Settimeoffintegrationstatusrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Settimeoffintegrationstatusrequest) MarshalJSON() ([]byte, error) {
    type Alias Settimeoffintegrationstatusrequest

    if SettimeoffintegrationstatusrequestMarshalled {
        return []byte("{}"), nil
    }
    SettimeoffintegrationstatusrequestMarshalled = true

    return json.Marshal(&struct {
        
        IntegrationStatus string `json:"integrationStatus"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


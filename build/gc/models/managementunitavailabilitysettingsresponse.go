package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManagementunitavailabilitysettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManagementunitavailabilitysettingsresponseDud struct { 
    

}

// Managementunitavailabilitysettingsresponse
type Managementunitavailabilitysettingsresponse struct { 
    // Enabled - Indicates whether agent availability is enabled for the management unit
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Managementunitavailabilitysettingsresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Managementunitavailabilitysettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Managementunitavailabilitysettingsresponse

    if ManagementunitavailabilitysettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    ManagementunitavailabilitysettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


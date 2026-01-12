package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryavailabilitymanagementunitssettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryavailabilitymanagementunitssettingsresponseDud struct { 
    

}

// Queryavailabilitymanagementunitssettingsresponse
type Queryavailabilitymanagementunitssettingsresponse struct { 
    // ManagementUnits - List of unavailable times settings, per management unit
    ManagementUnits []Unavailabletimesmanagementunitsettings `json:"managementUnits"`

}

// String returns a JSON representation of the model
func (o *Queryavailabilitymanagementunitssettingsresponse) String() string {
     o.ManagementUnits = []Unavailabletimesmanagementunitsettings{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryavailabilitymanagementunitssettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Queryavailabilitymanagementunitssettingsresponse

    if QueryavailabilitymanagementunitssettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    QueryavailabilitymanagementunitssettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnits []Unavailabletimesmanagementunitsettings `json:"managementUnits"`
        *Alias
    }{

        
        ManagementUnits: []Unavailabletimesmanagementunitsettings{{}},
        

        Alias: (*Alias)(u),
    })
}


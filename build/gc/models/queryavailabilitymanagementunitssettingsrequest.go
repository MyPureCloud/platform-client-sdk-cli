package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryavailabilitymanagementunitssettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryavailabilitymanagementunitssettingsrequestDud struct { 
    

}

// Queryavailabilitymanagementunitssettingsrequest
type Queryavailabilitymanagementunitssettingsrequest struct { 
    // ManagementUnitIds - The IDs of the management units for which to fetch their configurations. The management units must be in the business unit specified in the request path
    ManagementUnitIds []string `json:"managementUnitIds"`

}

// String returns a JSON representation of the model
func (o *Queryavailabilitymanagementunitssettingsrequest) String() string {
     o.ManagementUnitIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryavailabilitymanagementunitssettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryavailabilitymanagementunitssettingsrequest

    if QueryavailabilitymanagementunitssettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryavailabilitymanagementunitssettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        *Alias
    }{

        
        ManagementUnitIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}


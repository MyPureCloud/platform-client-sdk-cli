package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatemanagementunitssettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatemanagementunitssettingsrequestDud struct { 
    

}

// Updatemanagementunitssettingsrequest
type Updatemanagementunitssettingsrequest struct { 
    // Enabled - Indicates whether agent availability is enabled for the management unit
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Updatemanagementunitssettingsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatemanagementunitssettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatemanagementunitssettingsrequest

    if UpdatemanagementunitssettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatemanagementunitssettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


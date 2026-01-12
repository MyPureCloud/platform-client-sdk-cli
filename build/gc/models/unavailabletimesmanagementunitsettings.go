package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnavailabletimesmanagementunitsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnavailabletimesmanagementunitsettingsDud struct { 
    


    

}

// Unavailabletimesmanagementunitsettings
type Unavailabletimesmanagementunitsettings struct { 
    // ManagementUnit - The management unit to which these settings apply
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // Enabled - Indicates whether agent availability is enabled for the management unit
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Unavailabletimesmanagementunitsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unavailabletimesmanagementunitsettings) MarshalJSON() ([]byte, error) {
    type Alias Unavailabletimesmanagementunitsettings

    if UnavailabletimesmanagementunitsettingsMarshalled {
        return []byte("{}"), nil
    }
    UnavailabletimesmanagementunitsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


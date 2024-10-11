package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalcontactssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalcontactssettingsDud struct { 
    


    

}

// Externalcontactssettings
type Externalcontactssettings struct { 
    // ManuallyAssignDivisionsToInteractions
    ManuallyAssignDivisionsToInteractions bool `json:"manuallyAssignDivisionsToInteractions"`


    // ManuallyAssignDivisionsToContacts
    ManuallyAssignDivisionsToContacts bool `json:"manuallyAssignDivisionsToContacts"`

}

// String returns a JSON representation of the model
func (o *Externalcontactssettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalcontactssettings) MarshalJSON() ([]byte, error) {
    type Alias Externalcontactssettings

    if ExternalcontactssettingsMarshalled {
        return []byte("{}"), nil
    }
    ExternalcontactssettingsMarshalled = true

    return json.Marshal(&struct {
        
        ManuallyAssignDivisionsToInteractions bool `json:"manuallyAssignDivisionsToInteractions"`
        
        ManuallyAssignDivisionsToContacts bool `json:"manuallyAssignDivisionsToContacts"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


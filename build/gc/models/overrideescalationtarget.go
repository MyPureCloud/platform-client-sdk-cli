package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OverrideescalationtargetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OverrideescalationtargetDud struct { 
    

}

// Overrideescalationtarget
type Overrideescalationtarget struct { 
    // IntegrationId - The ID of the integration.
    IntegrationId string `json:"integrationId"`

}

// String returns a JSON representation of the model
func (o *Overrideescalationtarget) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Overrideescalationtarget) MarshalJSON() ([]byte, error) {
    type Alias Overrideescalationtarget

    if OverrideescalationtargetMarshalled {
        return []byte("{}"), nil
    }
    OverrideescalationtargetMarshalled = true

    return json.Marshal(&struct {
        
        IntegrationId string `json:"integrationId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


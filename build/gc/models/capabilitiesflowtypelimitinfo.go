package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesflowtypelimitinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesflowtypelimitinfoDud struct { 
    MaxActionsPerFlow Capabilitieslimitcondition `json:"maxActionsPerFlow"`


    MaxRuntimeDurationInSeconds Capabilitieslimitcondition `json:"maxRuntimeDurationInSeconds"`

}

// Capabilitiesflowtypelimitinfo
type Capabilitiesflowtypelimitinfo struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Capabilitiesflowtypelimitinfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesflowtypelimitinfo) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesflowtypelimitinfo

    if CapabilitiesflowtypelimitinfoMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesflowtypelimitinfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


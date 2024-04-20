package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportedentitytypestatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportedentitytypestatusDud struct { 
    

}

// Supportedentitytypestatus
type Supportedentitytypestatus struct { 
    // ListSlotType - The configuration status of restricted lists
    ListSlotType string `json:"listSlotType"`

}

// String returns a JSON representation of the model
func (o *Supportedentitytypestatus) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportedentitytypestatus) MarshalJSON() ([]byte, error) {
    type Alias Supportedentitytypestatus

    if SupportedentitytypestatusMarshalled {
        return []byte("{}"), nil
    }
    SupportedentitytypestatusMarshalled = true

    return json.Marshal(&struct {
        
        ListSlotType string `json:"listSlotType"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


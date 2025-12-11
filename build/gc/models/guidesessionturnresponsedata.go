package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionturnresponsedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionturnresponsedataDud struct { 
    

}

// Guidesessionturnresponsedata
type Guidesessionturnresponsedata struct { 
    // Text - The text response content.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Guidesessionturnresponsedata) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionturnresponsedata) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionturnresponsedata

    if GuidesessionturnresponsedataMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionturnresponsedataMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


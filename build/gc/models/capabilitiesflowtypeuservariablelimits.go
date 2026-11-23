package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesflowtypeuservariablelimitsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesflowtypeuservariablelimitsDud struct { 
    MaxCount int `json:"maxCount"`


    MaxSize int `json:"maxSize"`

}

// Capabilitiesflowtypeuservariablelimits
type Capabilitiesflowtypeuservariablelimits struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Capabilitiesflowtypeuservariablelimits) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesflowtypeuservariablelimits) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesflowtypeuservariablelimits

    if CapabilitiesflowtypeuservariablelimitsMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesflowtypeuservariablelimitsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


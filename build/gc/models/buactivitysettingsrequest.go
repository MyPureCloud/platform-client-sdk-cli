package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuactivitysettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuactivitysettingsrequestDud struct { 
    

}

// Buactivitysettingsrequest
type Buactivitysettingsrequest struct { 
    // DefaultActivityCodeId - Default Activity Code ID settings
    DefaultActivityCodeId string `json:"defaultActivityCodeId"`

}

// String returns a JSON representation of the model
func (o *Buactivitysettingsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buactivitysettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Buactivitysettingsrequest

    if BuactivitysettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    BuactivitysettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        DefaultActivityCodeId string `json:"defaultActivityCodeId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


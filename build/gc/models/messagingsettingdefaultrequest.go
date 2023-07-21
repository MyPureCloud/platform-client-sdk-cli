package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingsettingdefaultrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingsettingdefaultrequestDud struct { 
    

}

// Messagingsettingdefaultrequest
type Messagingsettingdefaultrequest struct { 
    // SettingId - Messaging Setting ID to be used as the default for this Organization.
    SettingId string `json:"settingId"`

}

// String returns a JSON representation of the model
func (o *Messagingsettingdefaultrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingsettingdefaultrequest) MarshalJSON() ([]byte, error) {
    type Alias Messagingsettingdefaultrequest

    if MessagingsettingdefaultrequestMarshalled {
        return []byte("{}"), nil
    }
    MessagingsettingdefaultrequestMarshalled = true

    return json.Marshal(&struct {
        
        SettingId string `json:"settingId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


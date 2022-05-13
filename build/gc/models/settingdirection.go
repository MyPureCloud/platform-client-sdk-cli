package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SettingdirectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SettingdirectionDud struct { 
    


    

}

// Settingdirection
type Settingdirection struct { 
    // Inbound - Status for the Inbound Direction
    Inbound string `json:"inbound"`


    // Outbound - Status for the Outbound Direction
    Outbound string `json:"outbound"`

}

// String returns a JSON representation of the model
func (o *Settingdirection) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Settingdirection) MarshalJSON() ([]byte, error) {
    type Alias Settingdirection

    if SettingdirectionMarshalled {
        return []byte("{}"), nil
    }
    SettingdirectionMarshalled = true

    return json.Marshal(&struct {
        
        Inbound string `json:"inbound"`
        
        Outbound string `json:"outbound"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


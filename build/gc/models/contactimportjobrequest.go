package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportjobrequestDud struct { 
    

}

// Contactimportjobrequest
type Contactimportjobrequest struct { 
    // SettingsId - Settings id
    SettingsId string `json:"settingsId"`

}

// String returns a JSON representation of the model
func (o *Contactimportjobrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactimportjobrequest

    if ContactimportjobrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactimportjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        SettingsId string `json:"settingsId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}


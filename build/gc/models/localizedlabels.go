package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocalizedlabelsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocalizedlabelsDud struct { 
    


    

}

// Localizedlabels - Contains localized labels used in messenger apps
type Localizedlabels struct { 
    // Key - Contains localized label key used in messenger homescreen and push notification. PushNotificationTitle and PushNotificationBody keys are required when notifications are enabled.
    Key string `json:"key"`


    // Value - Contains localized label value used in messenger homescreen and push notification
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Localizedlabels) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Localizedlabels) MarshalJSON() ([]byte, error) {
    type Alias Localizedlabels

    if LocalizedlabelsMarshalled {
        return []byte("{}"), nil
    }
    LocalizedlabelsMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


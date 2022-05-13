package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediasettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediasettingDud struct { 
    


    

}

// Mediasetting
type Mediasetting struct { 
    // AlertingTimeoutSeconds
    AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`


    // ServiceLevel
    ServiceLevel Servicelevel `json:"serviceLevel"`

}

// String returns a JSON representation of the model
func (o *Mediasetting) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediasetting) MarshalJSON() ([]byte, error) {
    type Alias Mediasetting

    if MediasettingMarshalled {
        return []byte("{}"), nil
    }
    MediasettingMarshalled = true

    return json.Marshal(&struct {
        
        AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`
        
        ServiceLevel Servicelevel `json:"serviceLevel"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


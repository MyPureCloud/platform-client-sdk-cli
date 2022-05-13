package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationstatusinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationstatusinfoDud struct { 
    Code string `json:"code"`


    Effective string `json:"effective"`


    Detail Messageinfo `json:"detail"`


    LastUpdated time.Time `json:"lastUpdated"`

}

// Integrationstatusinfo - Status information for an Integration.
type Integrationstatusinfo struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Integrationstatusinfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationstatusinfo) MarshalJSON() ([]byte, error) {
    type Alias Integrationstatusinfo

    if IntegrationstatusinfoMarshalled {
        return []byte("{}"), nil
    }
    IntegrationstatusinfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


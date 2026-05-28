package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaleventchangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaleventchangeDud struct { 
    ChangeCategory string `json:"changeCategory"`


    SchemaId string `json:"schemaId"`


    EventName string `json:"eventName"`


    DateDetected time.Time `json:"dateDetected"`


    SystemStatus string `json:"systemStatus"`


    ErrorCode string `json:"errorCode"`


    ErrorDescription string `json:"errorDescription"`

}

// Externaleventchange - A change in an external event definition
type Externaleventchange struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Externaleventchange) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externaleventchange) MarshalJSON() ([]byte, error) {
    type Alias Externaleventchange

    if ExternaleventchangeMarshalled {
        return []byte("{}"), nil
    }
    ExternaleventchangeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


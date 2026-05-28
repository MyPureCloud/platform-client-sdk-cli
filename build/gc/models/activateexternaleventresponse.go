package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivateexternaleventresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivateexternaleventresponseDud struct { 
    SchemaId string `json:"schemaId"`


    EventName string `json:"eventName"`


    DisplayName string `json:"displayName"`


    Rank int `json:"rank"`


    ActivationStatus string `json:"activationStatus"`

}

// Activateexternaleventresponse - Response for activation of an external event
type Activateexternaleventresponse struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Activateexternaleventresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activateexternaleventresponse) MarshalJSON() ([]byte, error) {
    type Alias Activateexternaleventresponse

    if ActivateexternaleventresponseMarshalled {
        return []byte("{}"), nil
    }
    ActivateexternaleventresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateexternaleventresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateexternaleventresponseDud struct { 
    SchemaId string `json:"schemaId"`


    EventName string `json:"eventName"`


    DisplayName string `json:"displayName"`


    Rank int `json:"rank"`


    ActivationStatus string `json:"activationStatus"`

}

// Updateexternaleventresponse - Response for updating an external event definition
type Updateexternaleventresponse struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Updateexternaleventresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateexternaleventresponse) MarshalJSON() ([]byte, error) {
    type Alias Updateexternaleventresponse

    if UpdateexternaleventresponseMarshalled {
        return []byte("{}"), nil
    }
    UpdateexternaleventresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaleventsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaleventsummaryDud struct { 
    SchemaId string `json:"schemaId"`


    EventName string `json:"eventName"`


    DisplayName string `json:"displayName"`


    Rank int `json:"rank"`


    ActivationStatus string `json:"activationStatus"`


    SystemStatus string `json:"systemStatus"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    DateFirstActivated time.Time `json:"dateFirstActivated"`

}

// Externaleventsummary - Summary of an external event definition
type Externaleventsummary struct { 
    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Externaleventsummary) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externaleventsummary) MarshalJSON() ([]byte, error) {
    type Alias Externaleventsummary

    if ExternaleventsummaryMarshalled {
        return []byte("{}"), nil
    }
    ExternaleventsummaryMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


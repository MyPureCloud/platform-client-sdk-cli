package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetexternaleventsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetexternaleventsresponseDud struct { 
    

}

// Getexternaleventsresponse - Response for listing external events
type Getexternaleventsresponse struct { 
    // Entities
    Entities []Externaleventsummary `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Getexternaleventsresponse) String() string {
     o.Entities = []Externaleventsummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getexternaleventsresponse) MarshalJSON() ([]byte, error) {
    type Alias Getexternaleventsresponse

    if GetexternaleventsresponseMarshalled {
        return []byte("{}"), nil
    }
    GetexternaleventsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Externaleventsummary `json:"entities"`
        *Alias
    }{

        
        Entities: []Externaleventsummary{{}},
        

        Alias: (*Alias)(u),
    })
}


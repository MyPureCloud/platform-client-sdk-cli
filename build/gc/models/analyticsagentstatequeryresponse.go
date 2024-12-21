package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsagentstatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsagentstatequeryresponseDud struct { 
    

}

// Analyticsagentstatequeryresponse
type Analyticsagentstatequeryresponse struct { 
    // Entities - List of agents
    Entities []Analyticsagentstateagentresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Analyticsagentstatequeryresponse) String() string {
     o.Entities = []Analyticsagentstateagentresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsagentstatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsagentstatequeryresponse

    if AnalyticsagentstatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsagentstatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Analyticsagentstateagentresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Analyticsagentstateagentresponse{{}},
        

        Alias: (*Alias)(u),
    })
}


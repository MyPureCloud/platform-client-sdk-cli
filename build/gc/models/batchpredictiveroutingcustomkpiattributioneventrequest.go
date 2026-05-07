package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchpredictiveroutingcustomkpiattributioneventrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchpredictiveroutingcustomkpiattributioneventrequestDud struct { 
    

}

// Batchpredictiveroutingcustomkpiattributioneventrequest - A maximum of 100 events are allowed per request
type Batchpredictiveroutingcustomkpiattributioneventrequest struct { 
    // CustomKpiAttributionEvents - PredictiveRoutingCustomKpiAttributionEvent events for this batch
    CustomKpiAttributionEvents []Predictiveroutingcustomkpiattributionevent `json:"customKpiAttributionEvents"`

}

// String returns a JSON representation of the model
func (o *Batchpredictiveroutingcustomkpiattributioneventrequest) String() string {
     o.CustomKpiAttributionEvents = []Predictiveroutingcustomkpiattributionevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchpredictiveroutingcustomkpiattributioneventrequest) MarshalJSON() ([]byte, error) {
    type Alias Batchpredictiveroutingcustomkpiattributioneventrequest

    if BatchpredictiveroutingcustomkpiattributioneventrequestMarshalled {
        return []byte("{}"), nil
    }
    BatchpredictiveroutingcustomkpiattributioneventrequestMarshalled = true

    return json.Marshal(&struct {
        
        CustomKpiAttributionEvents []Predictiveroutingcustomkpiattributionevent `json:"customKpiAttributionEvents"`
        *Alias
    }{

        
        CustomKpiAttributionEvents: []Predictiveroutingcustomkpiattributionevent{{}},
        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventaggregatesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventaggregatesresponseDud struct { 
    


    

}

// Eventaggregatesresponse
type Eventaggregatesresponse struct { 
    // Interval - Interval for returned aggregates. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // EventDefinitionAggregates - Aggregates by event definition
    EventDefinitionAggregates []Eventdefinitionaggregates `json:"eventDefinitionAggregates"`

}

// String returns a JSON representation of the model
func (o *Eventaggregatesresponse) String() string {
    
     o.EventDefinitionAggregates = []Eventdefinitionaggregates{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventaggregatesresponse) MarshalJSON() ([]byte, error) {
    type Alias Eventaggregatesresponse

    if EventaggregatesresponseMarshalled {
        return []byte("{}"), nil
    }
    EventaggregatesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        EventDefinitionAggregates []Eventdefinitionaggregates `json:"eventDefinitionAggregates"`
        *Alias
    }{

        


        
        EventDefinitionAggregates: []Eventdefinitionaggregates{{}},
        

        Alias: (*Alias)(u),
    })
}


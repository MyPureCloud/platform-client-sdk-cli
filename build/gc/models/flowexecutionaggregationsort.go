package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowexecutionaggregationsortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowexecutionaggregationsortDud struct { 
    


    

}

// Flowexecutionaggregationsort
type Flowexecutionaggregationsort struct { 
    // Name - Name of the metric used for sorting values.
    Name string `json:"name"`


    // Function - Aggregation function used for the sort metric.
    Function string `json:"function"`

}

// String returns a JSON representation of the model
func (o *Flowexecutionaggregationsort) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowexecutionaggregationsort) MarshalJSON() ([]byte, error) {
    type Alias Flowexecutionaggregationsort

    if FlowexecutionaggregationsortMarshalled {
        return []byte("{}"), nil
    }
    FlowexecutionaggregationsortMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Function string `json:"function"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


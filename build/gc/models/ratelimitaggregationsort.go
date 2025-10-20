package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RatelimitaggregationsortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RatelimitaggregationsortDud struct { 
    


    

}

// Ratelimitaggregationsort
type Ratelimitaggregationsort struct { 
    // Name - Name of the metric used for sorting values.
    Name string `json:"name"`


    // Function - Aggregation function used for the sort metric.
    Function string `json:"function"`

}

// String returns a JSON representation of the model
func (o *Ratelimitaggregationsort) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ratelimitaggregationsort) MarshalJSON() ([]byte, error) {
    type Alias Ratelimitaggregationsort

    if RatelimitaggregationsortMarshalled {
        return []byte("{}"), nil
    }
    RatelimitaggregationsortMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Function string `json:"function"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowaggregationsortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowaggregationsortDud struct { 
    


    

}

// Flowaggregationsort
type Flowaggregationsort struct { 
    // Name - Name of the metric used for sorting values.
    Name string `json:"name"`


    // Function - Aggregation function used for the sort metric.
    Function string `json:"function"`

}

// String returns a JSON representation of the model
func (o *Flowaggregationsort) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowaggregationsort) MarshalJSON() ([]byte, error) {
    type Alias Flowaggregationsort

    if FlowaggregationsortMarshalled {
        return []byte("{}"), nil
    }
    FlowaggregationsortMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Function string `json:"function"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


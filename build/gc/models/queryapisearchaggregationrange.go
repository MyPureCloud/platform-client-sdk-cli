package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryapisearchaggregationrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryapisearchaggregationrangeDud struct { 
    


    

}

// Queryapisearchaggregationrange
type Queryapisearchaggregationrange struct { 
    // To
    To float64 `json:"to"`


    // From
    From float64 `json:"from"`

}

// String returns a JSON representation of the model
func (o *Queryapisearchaggregationrange) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryapisearchaggregationrange) MarshalJSON() ([]byte, error) {
    type Alias Queryapisearchaggregationrange

    if QueryapisearchaggregationrangeMarshalled {
        return []byte("{}"), nil
    }
    QueryapisearchaggregationrangeMarshalled = true

    return json.Marshal(&struct {
        
        To float64 `json:"to"`
        
        From float64 `json:"from"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


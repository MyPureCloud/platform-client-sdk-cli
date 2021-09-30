package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregationrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregationrangeDud struct { 
    


    

}

// Aggregationrange
type Aggregationrange struct { 
    // Gte - Greater than or equal to
    Gte float32 `json:"gte"`


    // Lt - Less than
    Lt float32 `json:"lt"`

}

// String returns a JSON representation of the model
func (o *Aggregationrange) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregationrange) MarshalJSON() ([]byte, error) {
    type Alias Aggregationrange

    if AggregationrangeMarshalled {
        return []byte("{}"), nil
    }
    AggregationrangeMarshalled = true

    return json.Marshal(&struct { 
        Gte float32 `json:"gte"`
        
        Lt float32 `json:"lt"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


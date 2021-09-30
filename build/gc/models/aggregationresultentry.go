package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregationresultentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregationresultentryDud struct { 
    


    


    


    

}

// Aggregationresultentry
type Aggregationresultentry struct { 
    // Count
    Count int `json:"count"`


    // Value - For termFrequency aggregations
    Value string `json:"value"`


    // Gte - For numericRange aggregations
    Gte float32 `json:"gte"`


    // Lt - For numericRange aggregations
    Lt float32 `json:"lt"`

}

// String returns a JSON representation of the model
func (o *Aggregationresultentry) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregationresultentry) MarshalJSON() ([]byte, error) {
    type Alias Aggregationresultentry

    if AggregationresultentryMarshalled {
        return []byte("{}"), nil
    }
    AggregationresultentryMarshalled = true

    return json.Marshal(&struct { 
        Count int `json:"count"`
        
        Value string `json:"value"`
        
        Gte float32 `json:"gte"`
        
        Lt float32 `json:"lt"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


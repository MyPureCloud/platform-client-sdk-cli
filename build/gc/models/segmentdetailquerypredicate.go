package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentdetailquerypredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentdetailquerypredicateDud struct { 
    


    


    


    


    


    


    


    

}

// Segmentdetailquerypredicate
type Segmentdetailquerypredicate struct { 
    // VarType - Optional type, can usually be inferred
    VarType string `json:"type"`


    // Dimension - Left hand side for dimension predicates
    Dimension string `json:"dimension"`


    // PropertyType - Left hand side for property predicates
    PropertyType string `json:"propertyType"`


    // Property - Left hand side for property predicates
    Property string `json:"property"`


    // Metric - Left hand side for metric predicates
    Metric string `json:"metric"`


    // Operator - Optional operator, default is matches
    Operator string `json:"operator"`


    // Value - Right hand side for dimension, metric, or property predicates
    Value string `json:"value"`


    // VarRange - Right hand side for dimension, metric, or property predicates
    VarRange Numericrange `json:"range"`

}

// String returns a JSON representation of the model
func (o *Segmentdetailquerypredicate) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentdetailquerypredicate) MarshalJSON() ([]byte, error) {
    type Alias Segmentdetailquerypredicate

    if SegmentdetailquerypredicateMarshalled {
        return []byte("{}"), nil
    }
    SegmentdetailquerypredicateMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Dimension string `json:"dimension"`
        
        PropertyType string `json:"propertyType"`
        
        Property string `json:"property"`
        
        Metric string `json:"metric"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        
        VarRange Numericrange `json:"range"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


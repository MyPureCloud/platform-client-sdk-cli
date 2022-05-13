package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsqueryaggregationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsqueryaggregationDud struct { 
    


    


    


    


    

}

// Analyticsqueryaggregation
type Analyticsqueryaggregation struct { 
    // VarType - Optional type, can usually be inferred
    VarType string `json:"type"`


    // Dimension - For use with termFrequency aggregations
    Dimension string `json:"dimension"`


    // Metric - For use with numericRange aggregations
    Metric string `json:"metric"`


    // Size - For use with termFrequency aggregations
    Size int `json:"size"`


    // Ranges - For use with numericRange aggregations
    Ranges []Aggregationrange `json:"ranges"`

}

// String returns a JSON representation of the model
func (o *Analyticsqueryaggregation) String() string {
    
    
    
    
     o.Ranges = []Aggregationrange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsqueryaggregation) MarshalJSON() ([]byte, error) {
    type Alias Analyticsqueryaggregation

    if AnalyticsqueryaggregationMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsqueryaggregationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Dimension string `json:"dimension"`
        
        Metric string `json:"metric"`
        
        Size int `json:"size"`
        
        Ranges []Aggregationrange `json:"ranges"`
        *Alias
    }{

        


        


        


        


        
        Ranges: []Aggregationrange{{}},
        

        Alias: (*Alias)(u),
    })
}


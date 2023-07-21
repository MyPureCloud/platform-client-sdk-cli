package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommonrulepredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommonrulepredicateDud struct { 
    


    


    


    


    


    


    


    

}

// Commonrulepredicate
type Commonrulepredicate struct { 
    // MetricType - The type of metric being evaluated.
    MetricType string `json:"metricType"`


    // MetricValueType - The type of metric value being evaluated.
    MetricValueType string `json:"metricValueType"`


    // ComparisonOperator - The comparison operator being performed on the metric.
    ComparisonOperator string `json:"comparisonOperator"`


    // Value - The value the metric will be compared to.
    Value float64 `json:"value"`


    // Status - The status of the entity corresponding to the metric.
    Status string `json:"status"`


    // Entity - The entity whose metric is being represented.
    Entity Commonrulepredicateentity `json:"entity"`


    // MediaType - The media type of the conversation the metric describes.
    MediaType string `json:"mediaType"`


    // Metric - The metric being evaluated.
    Metric string `json:"metric"`

}

// String returns a JSON representation of the model
func (o *Commonrulepredicate) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commonrulepredicate) MarshalJSON() ([]byte, error) {
    type Alias Commonrulepredicate

    if CommonrulepredicateMarshalled {
        return []byte("{}"), nil
    }
    CommonrulepredicateMarshalled = true

    return json.Marshal(&struct {
        
        MetricType string `json:"metricType"`
        
        MetricValueType string `json:"metricValueType"`
        
        ComparisonOperator string `json:"comparisonOperator"`
        
        Value float64 `json:"value"`
        
        Status string `json:"status"`
        
        Entity Commonrulepredicateentity `json:"entity"`
        
        MediaType string `json:"mediaType"`
        
        Metric string `json:"metric"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


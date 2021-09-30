package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetricDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    LinkedMetric Addressableentityref `json:"linkedMetric"`


    DateCreated int `json:"dateCreated"`


    DateUnlinked time.Time `json:"dateUnlinked"`


    SourcePerformanceProfile Performanceprofile `json:"sourcePerformanceProfile"`


    SelfUri string `json:"selfUri"`

}

// Metric
type Metric struct { 
    


    // Name - The name of this metric
    Name string `json:"name"`


    // MetricDefinitionId - The id of associated metric definition
    MetricDefinitionId string `json:"metricDefinitionId"`


    // ExternalMetricDefinitionId - The id of associated external metric definition
    ExternalMetricDefinitionId string `json:"externalMetricDefinitionId"`


    // Objective - Associated objective for this metric
    Objective Objective `json:"objective"`


    // PerformanceProfileId - Performance profile id of this metric
    PerformanceProfileId string `json:"performanceProfileId"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Metric) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metric) MarshalJSON() ([]byte, error) {
    type Alias Metric

    if MetricMarshalled {
        return []byte("{}"), nil
    }
    MetricMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        MetricDefinitionId string `json:"metricDefinitionId"`
        
        ExternalMetricDefinitionId string `json:"externalMetricDefinitionId"`
        
        Objective Objective `json:"objective"`
        
        PerformanceProfileId string `json:"performanceProfileId"`
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


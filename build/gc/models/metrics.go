package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetricsDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    LinkedMetric Addressableentityref `json:"linkedMetric"`


    DateCreated time.Time `json:"dateCreated"`


    DateUnlinked time.Time `json:"dateUnlinked"`


    SourcePerformanceProfile Performanceprofile `json:"sourcePerformanceProfile"`


    UnitDefinition string `json:"unitDefinition"`


    Precision int `json:"precision"`


    SelfUri string `json:"selfUri"`

}

// Metrics
type Metrics struct { 
    


    // Name
    Name string `json:"name"`


    // Order - The order of metric within a performance profile
    Order int `json:"order"`


    // MetricDefinitionName - The name of associated metric definition
    MetricDefinitionName string `json:"metricDefinitionName"`


    // MetricDefinitionId - The id of associated metric definition
    MetricDefinitionId string `json:"metricDefinitionId"`


    // ExternalMetricDefinitionId - The id of associated external metric definition
    ExternalMetricDefinitionId string `json:"externalMetricDefinitionId"`


    // UnitType - Corresponding unit type for this metric
    UnitType string `json:"unitType"`


    // Enabled - A flag for whether this metric is enabled for a performance profile
    Enabled bool `json:"enabled"`


    // TemplateName - The name of associated objective template
    TemplateName string `json:"templateName"`


    // MaxPoints - Achievable maximum points for this metric
    MaxPoints int `json:"maxPoints"`


    // PerformanceProfileId - Performance profile id of this metric
    PerformanceProfileId string `json:"performanceProfileId"`


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Metrics) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metrics) MarshalJSON() ([]byte, error) {
    type Alias Metrics

    if MetricsMarshalled {
        return []byte("{}"), nil
    }
    MetricsMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Order int `json:"order"`
        
        MetricDefinitionName string `json:"metricDefinitionName"`
        
        MetricDefinitionId string `json:"metricDefinitionId"`
        
        ExternalMetricDefinitionId string `json:"externalMetricDefinitionId"`
        
        UnitType string `json:"unitType"`
        
        Enabled bool `json:"enabled"`
        
        TemplateName string `json:"templateName"`
        
        MaxPoints int `json:"maxPoints"`
        
        PerformanceProfileId string `json:"performanceProfileId"`
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


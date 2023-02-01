package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatemetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatemetricDud struct { 
    


    


    


    


    


    


    

}

// Createmetric
type Createmetric struct { 
    // MetricDefinitionId - The id of associated metric definition
    MetricDefinitionId string `json:"metricDefinitionId"`


    // ExternalMetricDefinitionId - The id of associated external metric definition
    ExternalMetricDefinitionId string `json:"externalMetricDefinitionId"`


    // Objective - Associated objective for this metric
    Objective Createobjective `json:"objective"`


    // PerformanceProfileId - Performance profile id of this metric
    PerformanceProfileId string `json:"performanceProfileId"`


    // Name - The name of this metric
    Name string `json:"name"`


    // Precision - The precision of the metric, must be between 0 and 5
    Precision int `json:"precision"`


    // TimeDisplayUnit - The time unit in which the metric should be displayed -- this parameter is ignored when displaying non-time values
    TimeDisplayUnit string `json:"timeDisplayUnit"`

}

// String returns a JSON representation of the model
func (o *Createmetric) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createmetric) MarshalJSON() ([]byte, error) {
    type Alias Createmetric

    if CreatemetricMarshalled {
        return []byte("{}"), nil
    }
    CreatemetricMarshalled = true

    return json.Marshal(&struct {
        
        MetricDefinitionId string `json:"metricDefinitionId"`
        
        ExternalMetricDefinitionId string `json:"externalMetricDefinitionId"`
        
        Objective Createobjective `json:"objective"`
        
        PerformanceProfileId string `json:"performanceProfileId"`
        
        Name string `json:"name"`
        
        Precision int `json:"precision"`
        
        TimeDisplayUnit string `json:"timeDisplayUnit"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


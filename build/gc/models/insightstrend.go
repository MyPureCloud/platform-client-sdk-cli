package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightstrendMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightstrendDud struct { 
    


    


    


    


    


    


    

}

// Insightstrend
type Insightstrend struct { 
    // PerformanceProfile - The performance profile
    PerformanceProfile Addressableentityref `json:"performanceProfile"`


    // Division - The division
    Division Divisionreference `json:"division"`


    // Granularity - Granularity
    Granularity string `json:"granularity"`


    // ComparativePeriod - The comparative period work day date range
    ComparativePeriod Workdayperiod `json:"comparativePeriod"`


    // PrimaryPeriod - The primary period work day date range
    PrimaryPeriod Workdayperiod `json:"primaryPeriod"`


    // Entities - The list of insights trend for each metric
    Entities []Insightstrendmetricitem `json:"entities"`


    // Total - The insights trend in total
    Total Insightstrendtotalitem `json:"total"`

}

// String returns a JSON representation of the model
func (o *Insightstrend) String() string {
    
    
    
    
    
     o.Entities = []Insightstrendmetricitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightstrend) MarshalJSON() ([]byte, error) {
    type Alias Insightstrend

    if InsightstrendMarshalled {
        return []byte("{}"), nil
    }
    InsightstrendMarshalled = true

    return json.Marshal(&struct {
        
        PerformanceProfile Addressableentityref `json:"performanceProfile"`
        
        Division Divisionreference `json:"division"`
        
        Granularity string `json:"granularity"`
        
        ComparativePeriod Workdayperiod `json:"comparativePeriod"`
        
        PrimaryPeriod Workdayperiod `json:"primaryPeriod"`
        
        Entities []Insightstrendmetricitem `json:"entities"`
        
        Total Insightstrendtotalitem `json:"total"`
        *Alias
    }{

        


        


        


        


        


        
        Entities: []Insightstrendmetricitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightsagentsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightsagentsDud struct { 
    


    


    


    


    


    

}

// Insightsagents
type Insightsagents struct { 
    // PerformanceProfile - The performance profile
    PerformanceProfile Addressableentityref `json:"performanceProfile"`


    // Division - The division
    Division Divisionreference `json:"division"`


    // Granularity - Granularity
    Granularity string `json:"granularity"`


    // DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    // DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEndWorkday time.Time `json:"dateEndWorkday"`


    // Entities - The list of insights agents
    Entities []Insightsagentitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Insightsagents) String() string {
    
    
    
    
    
     o.Entities = []Insightsagentitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightsagents) MarshalJSON() ([]byte, error) {
    type Alias Insightsagents

    if InsightsagentsMarshalled {
        return []byte("{}"), nil
    }
    InsightsagentsMarshalled = true

    return json.Marshal(&struct {
        
        PerformanceProfile Addressableentityref `json:"performanceProfile"`
        
        Division Divisionreference `json:"division"`
        
        Granularity string `json:"granularity"`
        
        DateStartWorkday time.Time `json:"dateStartWorkday"`
        
        DateEndWorkday time.Time `json:"dateEndWorkday"`
        
        Entities []Insightsagentitem `json:"entities"`
        *Alias
    }{

        


        


        


        


        


        
        Entities: []Insightsagentitem{{}},
        

        Alias: (*Alias)(u),
    })
}


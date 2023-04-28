package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightssummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightssummaryDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Insightssummary
type Insightssummary struct { 
    // Entities
    Entities []Insightssummaryuseritem `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


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


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Insightssummary) String() string {
     o.Entities = []Insightssummaryuseritem{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightssummary) MarshalJSON() ([]byte, error) {
    type Alias Insightssummary

    if InsightssummaryMarshalled {
        return []byte("{}"), nil
    }
    InsightssummaryMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Insightssummaryuseritem `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PerformanceProfile Addressableentityref `json:"performanceProfile"`
        
        Division Divisionreference `json:"division"`
        
        Granularity string `json:"granularity"`
        
        ComparativePeriod Workdayperiod `json:"comparativePeriod"`
        
        PrimaryPeriod Workdayperiod `json:"primaryPeriod"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Insightssummaryuseritem{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightsrankingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightsrankingsDud struct { 
    


    


    


    


    


    


    

}

// Insightsrankings
type Insightsrankings struct { 
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


    // Leaders - The leaders users
    Leaders []Insightssummaryuseritem `json:"leaders"`


    // Trailers - The trailing users
    Trailers []Insightssummaryuseritem `json:"trailers"`

}

// String returns a JSON representation of the model
func (o *Insightsrankings) String() string {
    
    
    
    
    
     o.Leaders = []Insightssummaryuseritem{{}} 
     o.Trailers = []Insightssummaryuseritem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightsrankings) MarshalJSON() ([]byte, error) {
    type Alias Insightsrankings

    if InsightsrankingsMarshalled {
        return []byte("{}"), nil
    }
    InsightsrankingsMarshalled = true

    return json.Marshal(&struct {
        
        PerformanceProfile Addressableentityref `json:"performanceProfile"`
        
        Division Divisionreference `json:"division"`
        
        Granularity string `json:"granularity"`
        
        ComparativePeriod Workdayperiod `json:"comparativePeriod"`
        
        PrimaryPeriod Workdayperiod `json:"primaryPeriod"`
        
        Leaders []Insightssummaryuseritem `json:"leaders"`
        
        Trailers []Insightssummaryuseritem `json:"trailers"`
        *Alias
    }{

        


        


        


        


        


        
        Leaders: []Insightssummaryuseritem{{}},
        


        
        Trailers: []Insightssummaryuseritem{{}},
        

        Alias: (*Alias)(u),
    })
}


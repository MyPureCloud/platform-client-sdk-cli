package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightsdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightsdetailsDud struct { 
    


    


    


    


    


    


    


    

}

// Insightsdetails
type Insightsdetails struct { 
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


    // User - The query user
    User Userreference `json:"user"`


    // MetricData - The list of insights data for each metric of the user
    MetricData []Insightsdetailsmetricitem `json:"metricData"`


    // OverallData - Overall insights data of the user
    OverallData Insightsdetailsoverallitem `json:"overallData"`

}

// String returns a JSON representation of the model
func (o *Insightsdetails) String() string {
    
    
    
    
    
    
     o.MetricData = []Insightsdetailsmetricitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightsdetails) MarshalJSON() ([]byte, error) {
    type Alias Insightsdetails

    if InsightsdetailsMarshalled {
        return []byte("{}"), nil
    }
    InsightsdetailsMarshalled = true

    return json.Marshal(&struct {
        
        PerformanceProfile Addressableentityref `json:"performanceProfile"`
        
        Division Divisionreference `json:"division"`
        
        Granularity string `json:"granularity"`
        
        ComparativePeriod Workdayperiod `json:"comparativePeriod"`
        
        PrimaryPeriod Workdayperiod `json:"primaryPeriod"`
        
        User Userreference `json:"user"`
        
        MetricData []Insightsdetailsmetricitem `json:"metricData"`
        
        OverallData Insightsdetailsoverallitem `json:"overallData"`
        *Alias
    }{

        


        


        


        


        


        


        
        MetricData: []Insightsdetailsmetricitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}


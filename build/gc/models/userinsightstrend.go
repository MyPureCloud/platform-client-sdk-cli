package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserinsightstrendMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserinsightstrendDud struct { 
    


    


    


    


    


    


    


    

}

// Userinsightstrend
type Userinsightstrend struct { 
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


    // User - The query user
    User Userreference `json:"user"`

}

// String returns a JSON representation of the model
func (o *Userinsightstrend) String() string {
    
    
    
    
    
     o.Entities = []Insightstrendmetricitem{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userinsightstrend) MarshalJSON() ([]byte, error) {
    type Alias Userinsightstrend

    if UserinsightstrendMarshalled {
        return []byte("{}"), nil
    }
    UserinsightstrendMarshalled = true

    return json.Marshal(&struct {
        
        PerformanceProfile Addressableentityref `json:"performanceProfile"`
        
        Division Divisionreference `json:"division"`
        
        Granularity string `json:"granularity"`
        
        ComparativePeriod Workdayperiod `json:"comparativePeriod"`
        
        PrimaryPeriod Workdayperiod `json:"primaryPeriod"`
        
        Entities []Insightstrendmetricitem `json:"entities"`
        
        Total Insightstrendtotalitem `json:"total"`
        
        User Userreference `json:"user"`
        *Alias
    }{

        


        


        


        


        


        
        Entities: []Insightstrendmetricitem{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}


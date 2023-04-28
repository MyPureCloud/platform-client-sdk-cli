package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightssummarymetricperiodpointsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightssummarymetricperiodpointsDud struct { 
    


    


    


    


    

}

// Insightssummarymetricperiodpoints
type Insightssummarymetricperiodpoints struct { 
    // Points - Points scored
    Points int `json:"points"`


    // MaxPoints - Max possible points
    MaxPoints int `json:"maxPoints"`


    // DataPointCount - Number of data points
    DataPointCount int `json:"dataPointCount"`


    // PercentOfGoal - Percentage of the goal
    PercentOfGoal float64 `json:"percentOfGoal"`


    // Value - Average value
    Value float64 `json:"value"`

}

// String returns a JSON representation of the model
func (o *Insightssummarymetricperiodpoints) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightssummarymetricperiodpoints) MarshalJSON() ([]byte, error) {
    type Alias Insightssummarymetricperiodpoints

    if InsightssummarymetricperiodpointsMarshalled {
        return []byte("{}"), nil
    }
    InsightssummarymetricperiodpointsMarshalled = true

    return json.Marshal(&struct {
        
        Points int `json:"points"`
        
        MaxPoints int `json:"maxPoints"`
        
        DataPointCount int `json:"dataPointCount"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        
        Value float64 `json:"value"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


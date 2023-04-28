package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightssummaryoverallperiodpointsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightssummaryoverallperiodpointsDud struct { 
    


    


    


    

}

// Insightssummaryoverallperiodpoints
type Insightssummaryoverallperiodpoints struct { 
    // Points - Points scored
    Points int `json:"points"`


    // MaxPoints - Max possible points
    MaxPoints int `json:"maxPoints"`


    // DataPointCount - Number of data points
    DataPointCount int `json:"dataPointCount"`


    // PercentOfGoal - Percentage of the goal
    PercentOfGoal float64 `json:"percentOfGoal"`

}

// String returns a JSON representation of the model
func (o *Insightssummaryoverallperiodpoints) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightssummaryoverallperiodpoints) MarshalJSON() ([]byte, error) {
    type Alias Insightssummaryoverallperiodpoints

    if InsightssummaryoverallperiodpointsMarshalled {
        return []byte("{}"), nil
    }
    InsightssummaryoverallperiodpointsMarshalled = true

    return json.Marshal(&struct {
        
        Points int `json:"points"`
        
        MaxPoints int `json:"maxPoints"`
        
        DataPointCount int `json:"dataPointCount"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


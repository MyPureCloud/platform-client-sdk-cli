package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightsdetailsoverallperiodpointsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightsdetailsoverallperiodpointsDud struct { 
    


    


    


    


    


    

}

// Insightsdetailsoverallperiodpoints
type Insightsdetailsoverallperiodpoints struct { 
    // Points - Points scored
    Points int `json:"points"`


    // MaxPoints - Max possible points
    MaxPoints int `json:"maxPoints"`


    // DataPointCount - Number of data points
    DataPointCount int `json:"dataPointCount"`


    // PercentOfGoal - Percentage of the goal
    PercentOfGoal float64 `json:"percentOfGoal"`


    // RankTotalPoints - The agent's rank in leader board for points on this metric
    RankTotalPoints int `json:"rankTotalPoints"`


    // RankPercentagePoints - The agent's rank in leader board for percentage on this metric
    RankPercentagePoints int `json:"rankPercentagePoints"`

}

// String returns a JSON representation of the model
func (o *Insightsdetailsoverallperiodpoints) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightsdetailsoverallperiodpoints) MarshalJSON() ([]byte, error) {
    type Alias Insightsdetailsoverallperiodpoints

    if InsightsdetailsoverallperiodpointsMarshalled {
        return []byte("{}"), nil
    }
    InsightsdetailsoverallperiodpointsMarshalled = true

    return json.Marshal(&struct {
        
        Points int `json:"points"`
        
        MaxPoints int `json:"maxPoints"`
        
        DataPointCount int `json:"dataPointCount"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        
        RankTotalPoints int `json:"rankTotalPoints"`
        
        RankPercentagePoints int `json:"rankPercentagePoints"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


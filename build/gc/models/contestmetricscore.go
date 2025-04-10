package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestmetricscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestmetricscoreDud struct { 
    


    


    


    

}

// Contestmetricscore
type Contestmetricscore struct { 
    // Metric - The gamification metric for the data
    Metric Addressableentityref `json:"metric"`


    // Score - The Contest Metric score
    Score float64 `json:"score"`


    // TotalPoints - The Contest Metric totalPoints
    TotalPoints float64 `json:"totalPoints"`


    // PercentOfGoal - The Contest Metric percentOfGoal
    PercentOfGoal float64 `json:"percentOfGoal"`

}

// String returns a JSON representation of the model
func (o *Contestmetricscore) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestmetricscore) MarshalJSON() ([]byte, error) {
    type Alias Contestmetricscore

    if ContestmetricscoreMarshalled {
        return []byte("{}"), nil
    }
    ContestmetricscoreMarshalled = true

    return json.Marshal(&struct {
        
        Metric Addressableentityref `json:"metric"`
        
        Score float64 `json:"score"`
        
        TotalPoints float64 `json:"totalPoints"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


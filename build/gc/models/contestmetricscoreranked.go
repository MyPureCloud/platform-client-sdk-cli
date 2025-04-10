package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestmetricscorerankedMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestmetricscorerankedDud struct { 
    


    


    


    


    

}

// Contestmetricscoreranked
type Contestmetricscoreranked struct { 
    // Metric - The gamification metric for the data
    Metric Addressableentityref `json:"metric"`


    // Score - The Contest Metric score
    Score float64 `json:"score"`


    // TotalPoints - The Contest Metric totalPoints
    TotalPoints float64 `json:"totalPoints"`


    // PercentOfGoal - The Contest Metric percentOfGoal
    PercentOfGoal float64 `json:"percentOfGoal"`


    // Rank - The Contest Score rank, a lower rank is better (1 is the best)
    Rank int `json:"rank"`

}

// String returns a JSON representation of the model
func (o *Contestmetricscoreranked) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestmetricscoreranked) MarshalJSON() ([]byte, error) {
    type Alias Contestmetricscoreranked

    if ContestmetricscorerankedMarshalled {
        return []byte("{}"), nil
    }
    ContestmetricscorerankedMarshalled = true

    return json.Marshal(&struct {
        
        Metric Addressableentityref `json:"metric"`
        
        Score float64 `json:"score"`
        
        TotalPoints float64 `json:"totalPoints"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        
        Rank int `json:"rank"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


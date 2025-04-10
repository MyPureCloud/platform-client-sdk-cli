package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscoreDud struct { 
    


    


    

}

// Contestscore
type Contestscore struct { 
    // Score - The Contest score
    Score float64 `json:"score"`


    // TotalPoints - The Contest totalPoints
    TotalPoints float64 `json:"totalPoints"`


    // PercentOfGoal - The Contest percentOfGoal
    PercentOfGoal float64 `json:"percentOfGoal"`

}

// String returns a JSON representation of the model
func (o *Contestscore) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscore) MarshalJSON() ([]byte, error) {
    type Alias Contestscore

    if ContestscoreMarshalled {
        return []byte("{}"), nil
    }
    ContestscoreMarshalled = true

    return json.Marshal(&struct {
        
        Score float64 `json:"score"`
        
        TotalPoints float64 `json:"totalPoints"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscorerankedMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscorerankedDud struct { 
    


    


    


    


    

}

// Contestscoreranked
type Contestscoreranked struct { 
    // Score - The Contest score
    Score float64 `json:"score"`


    // TotalPoints - The Contest totalPoints
    TotalPoints float64 `json:"totalPoints"`


    // PercentOfGoal - The Contest percentOfGoal
    PercentOfGoal float64 `json:"percentOfGoal"`


    // Rank - The Contest Score rank
    Rank int `json:"rank"`


    // Tier - The Contest Score tier
    Tier int `json:"tier"`

}

// String returns a JSON representation of the model
func (o *Contestscoreranked) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscoreranked) MarshalJSON() ([]byte, error) {
    type Alias Contestscoreranked

    if ContestscorerankedMarshalled {
        return []byte("{}"), nil
    }
    ContestscorerankedMarshalled = true

    return json.Marshal(&struct {
        
        Score float64 `json:"score"`
        
        TotalPoints float64 `json:"totalPoints"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        
        Rank int `json:"rank"`
        
        Tier int `json:"tier"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscoresagentsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscoresagentsDud struct { 
    


    


    


    

}

// Contestscoresagents
type Contestscoresagents struct { 
    // ContestScore - The Contest score
    ContestScore Contestscoreranked `json:"contestScore"`


    // MetricScores - The Contest metric scores
    MetricScores []Contestmetricscoreranked `json:"metricScores"`


    // Disqualified - Indicates whether an agent is disqualified or not
    Disqualified bool `json:"disqualified"`


    // User - The Contest Scores Leaderboard user
    User Userreference `json:"user"`

}

// String returns a JSON representation of the model
func (o *Contestscoresagents) String() string {
    
     o.MetricScores = []Contestmetricscoreranked{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscoresagents) MarshalJSON() ([]byte, error) {
    type Alias Contestscoresagents

    if ContestscoresagentsMarshalled {
        return []byte("{}"), nil
    }
    ContestscoresagentsMarshalled = true

    return json.Marshal(&struct {
        
        ContestScore Contestscoreranked `json:"contestScore"`
        
        MetricScores []Contestmetricscoreranked `json:"metricScores"`
        
        Disqualified bool `json:"disqualified"`
        
        User Userreference `json:"user"`
        *Alias
    }{

        


        
        MetricScores: []Contestmetricscoreranked{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}


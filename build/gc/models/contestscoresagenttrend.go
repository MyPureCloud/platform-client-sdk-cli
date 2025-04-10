package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscoresagenttrendMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscoresagenttrendDud struct { 
    


    


    


    

}

// Contestscoresagenttrend
type Contestscoresagenttrend struct { 
    // ContestScore - The Contest score
    ContestScore Contestscoreranked `json:"contestScore"`


    // MetricScores - The Contest metric scores
    MetricScores []Contestmetricscoreranked `json:"metricScores"`


    // Disqualified - Indicates whether an agent is disqualified or not
    Disqualified bool `json:"disqualified"`


    // DateWorkday - Workday of the contest scores leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateWorkday time.Time `json:"dateWorkday"`

}

// String returns a JSON representation of the model
func (o *Contestscoresagenttrend) String() string {
    
     o.MetricScores = []Contestmetricscoreranked{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscoresagenttrend) MarshalJSON() ([]byte, error) {
    type Alias Contestscoresagenttrend

    if ContestscoresagenttrendMarshalled {
        return []byte("{}"), nil
    }
    ContestscoresagenttrendMarshalled = true

    return json.Marshal(&struct {
        
        ContestScore Contestscoreranked `json:"contestScore"`
        
        MetricScores []Contestmetricscoreranked `json:"metricScores"`
        
        Disqualified bool `json:"disqualified"`
        
        DateWorkday time.Time `json:"dateWorkday"`
        *Alias
    }{

        


        
        MetricScores: []Contestmetricscoreranked{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}


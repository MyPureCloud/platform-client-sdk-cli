package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscoresgrouptrendMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscoresgrouptrendDud struct { 
    


    


    

}

// Contestscoresgrouptrend
type Contestscoresgrouptrend struct { 
    // DateWorkday - Workday of the contest score. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateWorkday time.Time `json:"dateWorkday"`


    // ContestScore - The Contest score
    ContestScore Contestscore `json:"contestScore"`


    // MetricScores - The Contest metric scores
    MetricScores []Contestmetricscore `json:"metricScores"`

}

// String returns a JSON representation of the model
func (o *Contestscoresgrouptrend) String() string {
    
    
     o.MetricScores = []Contestmetricscore{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscoresgrouptrend) MarshalJSON() ([]byte, error) {
    type Alias Contestscoresgrouptrend

    if ContestscoresgrouptrendMarshalled {
        return []byte("{}"), nil
    }
    ContestscoresgrouptrendMarshalled = true

    return json.Marshal(&struct {
        
        DateWorkday time.Time `json:"dateWorkday"`
        
        ContestScore Contestscore `json:"contestScore"`
        
        MetricScores []Contestmetricscore `json:"metricScores"`
        *Alias
    }{

        


        


        
        MetricScores: []Contestmetricscore{{}},
        

        Alias: (*Alias)(u),
    })
}


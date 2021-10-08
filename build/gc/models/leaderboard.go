package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LeaderboardMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LeaderboardDud struct { 
    Division Division `json:"division"`


    Metric Addressableentityref `json:"metric"`


    DateStartWorkday time.Time `json:"dateStartWorkday"`


    DateEndWorkday time.Time `json:"dateEndWorkday"`


    Leaders []Leaderboarditem `json:"leaders"`


    UserRank Leaderboarditem `json:"userRank"`


    PerformanceProfile Addressableentityref `json:"performanceProfile"`

}

// Leaderboard
type Leaderboard struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Leaderboard) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Leaderboard) MarshalJSON() ([]byte, error) {
    type Alias Leaderboard

    if LeaderboardMarshalled {
        return []byte("{}"), nil
    }
    LeaderboardMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


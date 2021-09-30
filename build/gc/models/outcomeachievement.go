package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeachievementMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeachievementDud struct { 
    


    

}

// Outcomeachievement
type Outcomeachievement struct { 
    // Outcome - The outcome that was achieved.
    Outcome Achievedoutcome `json:"outcome"`


    // AchievedDate - Timestamp indicating when the outcome was achieved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    AchievedDate time.Time `json:"achievedDate"`

}

// String returns a JSON representation of the model
func (o *Outcomeachievement) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeachievement) MarshalJSON() ([]byte, error) {
    type Alias Outcomeachievement

    if OutcomeachievementMarshalled {
        return []byte("{}"), nil
    }
    OutcomeachievementMarshalled = true

    return json.Marshal(&struct { 
        Outcome Achievedoutcome `json:"outcome"`
        
        AchievedDate time.Time `json:"achievedDate"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


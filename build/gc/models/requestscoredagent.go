package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestscoredagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestscoredagentDud struct { 
    


    

}

// Requestscoredagent
type Requestscoredagent struct { 
    // Id - Agent's user ID
    Id string `json:"id"`


    // Score - Agent's score for the current conversation, from 0 - 100, higher being better
    Score int `json:"score"`

}

// String returns a JSON representation of the model
func (o *Requestscoredagent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestscoredagent) MarshalJSON() ([]byte, error) {
    type Alias Requestscoredagent

    if RequestscoredagentMarshalled {
        return []byte("{}"), nil
    }
    RequestscoredagentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Score int `json:"score"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemscoredagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemscoredagentDud struct { 
    


    

}

// Workitemscoredagent
type Workitemscoredagent struct { 
    // Agent - The agent
    Agent Userreference `json:"agent"`


    // Score - Agent's score for the workitem, from 0 - 100, higher being better
    Score int `json:"score"`

}

// String returns a JSON representation of the model
func (o *Workitemscoredagent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemscoredagent) MarshalJSON() ([]byte, error) {
    type Alias Workitemscoredagent

    if WorkitemscoredagentMarshalled {
        return []byte("{}"), nil
    }
    WorkitemscoredagentMarshalled = true

    return json.Marshal(&struct {
        
        Agent Userreference `json:"agent"`
        
        Score int `json:"score"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


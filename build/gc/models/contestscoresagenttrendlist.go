package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscoresagenttrendlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscoresagenttrendlistDud struct { 
    


    

}

// Contestscoresagenttrendlist
type Contestscoresagenttrendlist struct { 
    // Entities
    Entities []Contestscoresagenttrend `json:"entities"`


    // User - The Contest Scores Leaderboard user
    User Userreference `json:"user"`

}

// String returns a JSON representation of the model
func (o *Contestscoresagenttrendlist) String() string {
     o.Entities = []Contestscoresagenttrend{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscoresagenttrendlist) MarshalJSON() ([]byte, error) {
    type Alias Contestscoresagenttrendlist

    if ContestscoresagenttrendlistMarshalled {
        return []byte("{}"), nil
    }
    ContestscoresagenttrendlistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contestscoresagenttrend `json:"entities"`
        
        User Userreference `json:"user"`
        *Alias
    }{

        
        Entities: []Contestscoresagenttrend{{}},
        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemscoredagentdeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemscoredagentdeltaDud struct { 
    


    

}

// Workitemscoredagentdelta
type Workitemscoredagentdelta struct { 
    // Id
    Id string `json:"id"`


    // Score
    Score int `json:"score"`

}

// String returns a JSON representation of the model
func (o *Workitemscoredagentdelta) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemscoredagentdelta) MarshalJSON() ([]byte, error) {
    type Alias Workitemscoredagentdelta

    if WorkitemscoredagentdeltaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemscoredagentdeltaMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Score int `json:"score"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


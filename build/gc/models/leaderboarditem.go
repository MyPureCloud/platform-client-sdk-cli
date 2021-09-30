package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LeaderboarditemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LeaderboarditemDud struct { 
    User Userreference `json:"user"`


    Rank int `json:"rank"`


    Points int `json:"points"`

}

// Leaderboarditem
type Leaderboarditem struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Leaderboarditem) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Leaderboarditem) MarshalJSON() ([]byte, error) {
    type Alias Leaderboarditem

    if LeaderboarditemMarshalled {
        return []byte("{}"), nil
    }
    LeaderboarditemMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


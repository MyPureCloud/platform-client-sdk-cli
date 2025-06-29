package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestuserrankMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestuserrankDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Contestuserrank
type Contestuserrank struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Rank - The user's rank in contest, a lower rank is better (1 is the best)
    Rank int `json:"rank"`


    // Score - The user's contest score
    Score float64 `json:"score"`


    

}

// String returns a JSON representation of the model
func (o *Contestuserrank) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestuserrank) MarshalJSON() ([]byte, error) {
    type Alias Contestuserrank

    if ContestuserrankMarshalled {
        return []byte("{}"), nil
    }
    ContestuserrankMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Rank int `json:"rank"`
        
        Score float64 `json:"score"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


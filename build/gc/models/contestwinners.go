package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestwinnersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestwinnersDud struct { 
    


    


    

}

// Contestwinners
type Contestwinners struct { 
    // Tier - The Contest Winner tier
    Tier int `json:"tier"`


    // WinnersCount - The number of Contest Winners in a tier
    WinnersCount int `json:"winnersCount"`


    // Users - The Contest Winner users at the tier
    Users []Contestuserrank `json:"users"`

}

// String returns a JSON representation of the model
func (o *Contestwinners) String() string {
    
    
     o.Users = []Contestuserrank{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestwinners) MarshalJSON() ([]byte, error) {
    type Alias Contestwinners

    if ContestwinnersMarshalled {
        return []byte("{}"), nil
    }
    ContestwinnersMarshalled = true

    return json.Marshal(&struct {
        
        Tier int `json:"tier"`
        
        WinnersCount int `json:"winnersCount"`
        
        Users []Contestuserrank `json:"users"`
        *Alias
    }{

        


        


        
        Users: []Contestuserrank{{}},
        

        Alias: (*Alias)(u),
    })
}


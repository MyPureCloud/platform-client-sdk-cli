package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestdatawinnersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestdatawinnersDud struct { 
    


    


    


    

}

// Contestdatawinners
type Contestdatawinners struct { 
    // Tier - Tier of the winners
    Tier int `json:"tier"`


    // WinnersCount - Number of winners in this tier
    WinnersCount int `json:"winnersCount"`


    // ContestScore - Number of winners in this tier
    ContestScore Contestcompletedatascore `json:"contestScore"`


    // Users - List of users in this tier
    Users []Contestdatawinnersusers `json:"users"`

}

// String returns a JSON representation of the model
func (o *Contestdatawinners) String() string {
    
    
    
     o.Users = []Contestdatawinnersusers{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestdatawinners) MarshalJSON() ([]byte, error) {
    type Alias Contestdatawinners

    if ContestdatawinnersMarshalled {
        return []byte("{}"), nil
    }
    ContestdatawinnersMarshalled = true

    return json.Marshal(&struct {
        
        Tier int `json:"tier"`
        
        WinnersCount int `json:"winnersCount"`
        
        ContestScore Contestcompletedatascore `json:"contestScore"`
        
        Users []Contestdatawinnersusers `json:"users"`
        *Alias
    }{

        


        


        


        
        Users: []Contestdatawinnersusers{{}},
        

        Alias: (*Alias)(u),
    })
}


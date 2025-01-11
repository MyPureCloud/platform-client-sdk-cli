package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestdatawinnersusersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestdatawinnersusersDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Contestdatawinnersusers
type Contestdatawinnersusers struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Rank - The Contest user's rank
    Rank int `json:"rank"`


    

}

// String returns a JSON representation of the model
func (o *Contestdatawinnersusers) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestdatawinnersusers) MarshalJSON() ([]byte, error) {
    type Alias Contestdatawinnersusers

    if ContestdatawinnersusersMarshalled {
        return []byte("{}"), nil
    }
    ContestdatawinnersusersMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Rank int `json:"rank"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemscoredagentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemscoredagentrequestDud struct { 
    


    

}

// Workitemscoredagentrequest
type Workitemscoredagentrequest struct { 
    // Id - An agents ID. Must be a valid UUID.
    Id string `json:"id"`


    // Score - Agent's score for the workitem, from 0 - 100, higher being better
    Score int `json:"score"`

}

// String returns a JSON representation of the model
func (o *Workitemscoredagentrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemscoredagentrequest) MarshalJSON() ([]byte, error) {
    type Alias Workitemscoredagentrequest

    if WorkitemscoredagentrequestMarshalled {
        return []byte("{}"), nil
    }
    WorkitemscoredagentrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Score int `json:"score"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestdisqualifiedagentsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestdisqualifiedagentsDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Contestdisqualifiedagents
type Contestdisqualifiedagents struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Note - The disqualification explanation
    Note string `json:"note"`


    

}

// String returns a JSON representation of the model
func (o *Contestdisqualifiedagents) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestdisqualifiedagents) MarshalJSON() ([]byte, error) {
    type Alias Contestdisqualifiedagents

    if ContestdisqualifiedagentsMarshalled {
        return []byte("{}"), nil
    }
    ContestdisqualifiedagentsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Note string `json:"note"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


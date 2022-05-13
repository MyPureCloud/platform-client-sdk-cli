package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingskillMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingskillDud struct { 
    Id string `json:"id"`


    


    DateModified time.Time `json:"dateModified"`


    State string `json:"state"`


    Version string `json:"version"`


    SelfUri string `json:"selfUri"`

}

// Routingskill
type Routingskill struct { 
    


    // Name - The name of the skill.
    Name string `json:"name"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Routingskill) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingskill) MarshalJSON() ([]byte, error) {
    type Alias Routingskill

    if RoutingskillMarshalled {
        return []byte("{}"), nil
    }
    RoutingskillMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MoveagentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MoveagentresponseDud struct { 
    


    

}

// Moveagentresponse
type Moveagentresponse struct { 
    // User - The user associated with the move
    User Userreference `json:"user"`


    // Result - The result of the move
    Result string `json:"result"`

}

// String returns a JSON representation of the model
func (o *Moveagentresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Moveagentresponse) MarshalJSON() ([]byte, error) {
    type Alias Moveagentresponse

    if MoveagentresponseMarshalled {
        return []byte("{}"), nil
    }
    MoveagentresponseMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Result string `json:"result"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


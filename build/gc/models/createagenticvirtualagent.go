package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateagenticvirtualagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateagenticvirtualagentDud struct { 
    


    

}

// Createagenticvirtualagent
type Createagenticvirtualagent struct { 
    // Name - The name of the virtual agent.
    Name string `json:"name"`


    // ImageUri - The URI of the image for the virtual agent.
    ImageUri string `json:"imageUri"`

}

// String returns a JSON representation of the model
func (o *Createagenticvirtualagent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createagenticvirtualagent) MarshalJSON() ([]byte, error) {
    type Alias Createagenticvirtualagent

    if CreateagenticvirtualagentMarshalled {
        return []byte("{}"), nil
    }
    CreateagenticvirtualagentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ImageUri string `json:"imageUri"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


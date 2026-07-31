package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateagenticvirtualagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateagenticvirtualagentDud struct { 
    


    

}

// Updateagenticvirtualagent
type Updateagenticvirtualagent struct { 
    // Name - The name of the virtual agent.
    Name string `json:"name"`


    // ImageUri - The URI of the image for the virtual agent.
    ImageUri string `json:"imageUri"`

}

// String returns a JSON representation of the model
func (o *Updateagenticvirtualagent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateagenticvirtualagent) MarshalJSON() ([]byte, error) {
    type Alias Updateagenticvirtualagent

    if UpdateagenticvirtualagentMarshalled {
        return []byte("{}"), nil
    }
    UpdateagenticvirtualagentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ImageUri string `json:"imageUri"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


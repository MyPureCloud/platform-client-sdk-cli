package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Status string `json:"status"`


    LatestSavedVersion Agenticversionaddressableentity `json:"latestSavedVersion"`


    LatestProductionReadyVersion Agenticversionaddressableentity `json:"latestProductionReadyVersion"`


    


    SelfUri string `json:"selfUri"`

}

// Agenticvirtualagent
type Agenticvirtualagent struct { 
    


    // Name
    Name string `json:"name"`


    


    


    


    


    


    // ImageUri - The URI of the image for the virtual agent.
    ImageUri string `json:"imageUri"`


    

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagent) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagent

    if AgenticvirtualagentMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ImageUri string `json:"imageUri"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


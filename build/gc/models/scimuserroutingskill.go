package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimuserroutingskillMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimuserroutingskillDud struct { 
    


    

}

// Scimuserroutingskill - The routing skill assigned to a user.
type Scimuserroutingskill struct { 
    // Name - The case-sensitive name of a routing skill configured in Genesys Cloud.
    Name string `json:"name"`


    // Proficiency - A rating from 0.0 to 5.0 that indicates how adept an agent is at a particular skill. When \"Best available skills\" is enabled for a queue in Genesys Cloud, ACD interactions in that queue are routed to agents with higher proficiency ratings.
    Proficiency float64 `json:"proficiency"`

}

// String returns a JSON representation of the model
func (o *Scimuserroutingskill) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimuserroutingskill) MarshalJSON() ([]byte, error) {
    type Alias Scimuserroutingskill

    if ScimuserroutingskillMarshalled {
        return []byte("{}"), nil
    }
    ScimuserroutingskillMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Proficiency float64 `json:"proficiency"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


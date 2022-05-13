package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserroutingskillMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserroutingskillDud struct { 
    Id string `json:"id"`


    


    


    


    SkillUri string `json:"skillUri"`


    SelfUri string `json:"selfUri"`

}

// Userroutingskill - Represents an organization skill assigned to a user. When assigning to a user specify the organization skill id as the id.
type Userroutingskill struct { 
    


    // Name
    Name string `json:"name"`


    // Proficiency - A rating from 0.0 to 5.0 that indicates how adept an agent is at a particular skill. When \"Best available skills\" is enabled for a queue in Genesys Cloud, ACD interactions in that queue are routed to agents with higher proficiency ratings.
    Proficiency float64 `json:"proficiency"`


    // State - Activate or deactivate this routing skill.
    State string `json:"state"`


    


    

}

// String returns a JSON representation of the model
func (o *Userroutingskill) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userroutingskill) MarshalJSON() ([]byte, error) {
    type Alias Userroutingskill

    if UserroutingskillMarshalled {
        return []byte("{}"), nil
    }
    UserroutingskillMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Proficiency float64 `json:"proficiency"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


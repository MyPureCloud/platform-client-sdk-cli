package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserroutingskillpostMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserroutingskillpostDud struct { 
    


    


    SkillUri string `json:"skillUri"`


    SelfUri string `json:"selfUri"`

}

// Userroutingskillpost - Represents an organization skill assigned to a user. When assigning to a user specify the organization skill id as the id.
type Userroutingskillpost struct { 
    // Id - The id of the existing routing skill to add to the user
    Id string `json:"id"`


    // Proficiency - Proficiency is a rating from 0.0 to 5.0 on how competent an agent is for a particular skill. It is used when a queue is set to \"Best available skills\" mode to allow acd interactions to target agents with higher proficiency ratings.
    Proficiency float64 `json:"proficiency"`


    


    

}

// String returns a JSON representation of the model
func (o *Userroutingskillpost) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userroutingskillpost) MarshalJSON() ([]byte, error) {
    type Alias Userroutingskillpost

    if UserroutingskillpostMarshalled {
        return []byte("{}"), nil
    }
    UserroutingskillpostMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Proficiency float64 `json:"proficiency"`
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


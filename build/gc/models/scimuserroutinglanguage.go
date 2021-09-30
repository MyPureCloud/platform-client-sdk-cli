package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimuserroutinglanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimuserroutinglanguageDud struct { 
    


    

}

// Scimuserroutinglanguage - The routing language assigned to a user.
type Scimuserroutinglanguage struct { 
    // Name - The case-sensitive name of a routing language configured in Genesys Cloud.
    Name string `json:"name"`


    // Proficiency - A rating from 0.0 to 5.0 that indicates how fluent an agent is in a particular language. ACD interactions are routed to agents with higher proficiency ratings.
    Proficiency float64 `json:"proficiency"`

}

// String returns a JSON representation of the model
func (o *Scimuserroutinglanguage) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimuserroutinglanguage) MarshalJSON() ([]byte, error) {
    type Alias Scimuserroutinglanguage

    if ScimuserroutinglanguageMarshalled {
        return []byte("{}"), nil
    }
    ScimuserroutinglanguageMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Proficiency float64 `json:"proficiency"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


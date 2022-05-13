package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserroutinglanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserroutinglanguageDud struct { 
    Id string `json:"id"`


    


    


    


    LanguageUri string `json:"languageUri"`


    SelfUri string `json:"selfUri"`

}

// Userroutinglanguage - Represents an organization language assigned to a user. When assigning to a user specify the organization language id as the id.
type Userroutinglanguage struct { 
    


    // Name
    Name string `json:"name"`


    // Proficiency - A rating from 0.0 to 5.0 that indicates how fluent an agent is in a particular language. ACD interactions are routed to agents with higher proficiency ratings.
    Proficiency float64 `json:"proficiency"`


    // State - Activate or deactivate this routing language.
    State string `json:"state"`


    


    

}

// String returns a JSON representation of the model
func (o *Userroutinglanguage) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userroutinglanguage) MarshalJSON() ([]byte, error) {
    type Alias Userroutinglanguage

    if UserroutinglanguageMarshalled {
        return []byte("{}"), nil
    }
    UserroutinglanguageMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Proficiency float64 `json:"proficiency"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


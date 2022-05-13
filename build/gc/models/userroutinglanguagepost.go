package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserroutinglanguagepostMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserroutinglanguagepostDud struct { 
    


    


    LanguageUri string `json:"languageUri"`


    SelfUri string `json:"selfUri"`

}

// Userroutinglanguagepost - Represents an organization language assigned to a user. When assigning to a user specify the organization langauge id as the id.
type Userroutinglanguagepost struct { 
    // Id - The id of the existing routing language to add to the user
    Id string `json:"id"`


    // Proficiency - Proficiency is a rating from 0.0 to 5.0 on how competent an agent is for a particular language. It is used when a queue is set to \"Best available language\" mode to allow acd interactions to target agents with higher proficiency ratings.
    Proficiency float64 `json:"proficiency"`


    


    

}

// String returns a JSON representation of the model
func (o *Userroutinglanguagepost) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userroutinglanguagepost) MarshalJSON() ([]byte, error) {
    type Alias Userroutinglanguagepost

    if UserroutinglanguagepostMarshalled {
        return []byte("{}"), nil
    }
    UserroutinglanguagepostMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Proficiency float64 `json:"proficiency"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


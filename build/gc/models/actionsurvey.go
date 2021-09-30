package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionsurveyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionsurveyDud struct { 
    

}

// Actionsurvey
type Actionsurvey struct { 
    // Questions - Questions shown to the user.
    Questions []Journeysurveyquestion `json:"questions"`

}

// String returns a JSON representation of the model
func (o *Actionsurvey) String() string {
    
    
     o.Questions = []Journeysurveyquestion{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionsurvey) MarshalJSON() ([]byte, error) {
    type Alias Actionsurvey

    if ActionsurveyMarshalled {
        return []byte("{}"), nil
    }
    ActionsurveyMarshalled = true

    return json.Marshal(&struct { 
        Questions []Journeysurveyquestion `json:"questions"`
        
        *Alias
    }{
        

        
        Questions: []Journeysurveyquestion{{}},
        

        
        Alias: (*Alias)(u),
    })
}


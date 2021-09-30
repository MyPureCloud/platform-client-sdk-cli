package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchactionsurveyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchactionsurveyDud struct { 
    

}

// Patchactionsurvey
type Patchactionsurvey struct { 
    // Questions - Questions shown to the user.
    Questions []Patchsurveyquestion `json:"questions"`

}

// String returns a JSON representation of the model
func (o *Patchactionsurvey) String() string {
    
    
     o.Questions = []Patchsurveyquestion{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchactionsurvey) MarshalJSON() ([]byte, error) {
    type Alias Patchactionsurvey

    if PatchactionsurveyMarshalled {
        return []byte("{}"), nil
    }
    PatchactionsurveyMarshalled = true

    return json.Marshal(&struct { 
        Questions []Patchsurveyquestion `json:"questions"`
        
        *Alias
    }{
        

        
        Questions: []Patchsurveyquestion{{}},
        

        
        Alias: (*Alias)(u),
    })
}


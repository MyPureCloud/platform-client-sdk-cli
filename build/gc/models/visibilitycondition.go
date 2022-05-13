package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VisibilityconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VisibilityconditionDud struct { 
    


    

}

// Visibilitycondition
type Visibilitycondition struct { 
    // CombiningOperation
    CombiningOperation string `json:"combiningOperation"`


    // Predicates - A list of strings, each representing the location in the form of the Answer Option to depend on. In the format of \"/form/questionGroup/{questionGroupIndex}/question/{questionIndex}/answer/{answerIndex}\" or, to assume the current question group, \"../question/{questionIndex}/answer/{answerIndex}\". Note: Indexes are zero-based
    Predicates []interface{} `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Visibilitycondition) String() string {
    
     o.Predicates = []interface{}{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Visibilitycondition) MarshalJSON() ([]byte, error) {
    type Alias Visibilitycondition

    if VisibilityconditionMarshalled {
        return []byte("{}"), nil
    }
    VisibilityconditionMarshalled = true

    return json.Marshal(&struct {
        
        CombiningOperation string `json:"combiningOperation"`
        
        Predicates []interface{} `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []interface{}{},
        

        Alias: (*Alias)(u),
    })
}


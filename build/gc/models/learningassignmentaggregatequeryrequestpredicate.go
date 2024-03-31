package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentaggregatequeryrequestpredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentaggregatequeryrequestpredicateDud struct { 
    


    

}

// Learningassignmentaggregatequeryrequestpredicate
type Learningassignmentaggregatequeryrequestpredicate struct { 
    // Dimension - Each predicates specifies a dimension.
    Dimension string `json:"dimension"`


    // Value - Corresponding value for dimensions in predicates. If the dimension is type, Valid Values: Informational, AssessedContent, Assessment, External
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestpredicate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentaggregatequeryrequestpredicate) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentaggregatequeryrequestpredicate

    if LearningassignmentaggregatequeryrequestpredicateMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentaggregatequeryrequestpredicateMarshalled = true

    return json.Marshal(&struct {
        
        Dimension string `json:"dimension"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


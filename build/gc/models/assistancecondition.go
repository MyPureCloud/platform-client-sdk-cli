package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistanceconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistanceconditionDud struct { 
    


    

}

// Assistancecondition
type Assistancecondition struct { 
    // Operator - The operator for the assistance condition. The operator defines whether the listed topicIds should EXIST or NOTEXIST for the condition to be evaluated as true.
    Operator string `json:"operator"`


    // TopicIds - List of topicIds within the assistance condition which would be combined together using logical OR operator. Eg ( topicId_1 || topicId_2 ) .
    TopicIds []string `json:"topicIds"`

}

// String returns a JSON representation of the model
func (o *Assistancecondition) String() string {
    
     o.TopicIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistancecondition) MarshalJSON() ([]byte, error) {
    type Alias Assistancecondition

    if AssistanceconditionMarshalled {
        return []byte("{}"), nil
    }
    AssistanceconditionMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        TopicIds []string `json:"topicIds"`
        *Alias
    }{

        


        
        TopicIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}


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
    // Operator
    Operator string `json:"operator"`


    // TopicIds
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


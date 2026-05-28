package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillexpressionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillexpressionDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Skillexpression - A skill expression entity with ID, expression string (raw or normalized), and queue ID
type Skillexpression struct { 
    


    // Name
    Name string `json:"name"`


    // Expression - The skill expression string (raw or normalized, as requested)
    Expression string `json:"expression"`


    // QueueId - The queue ID where the expression is used
    QueueId string `json:"queueId"`


    

}

// String returns a JSON representation of the model
func (o *Skillexpression) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillexpression) MarshalJSON() ([]byte, error) {
    type Alias Skillexpression

    if SkillexpressionMarshalled {
        return []byte("{}"), nil
    }
    SkillexpressionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Expression string `json:"expression"`
        
        QueueId string `json:"queueId"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


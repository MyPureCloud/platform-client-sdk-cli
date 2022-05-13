package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulerulepartsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulerulepartsDud struct { 
    


    


    


    

}

// Learningmoduleruleparts
type Learningmoduleruleparts struct { 
    // Operation - The learning module rule operation
    Operation string `json:"operation"`


    // Selector - The learning module rule selector
    Selector string `json:"selector"`


    // Value - The value of rules
    Value []string `json:"value"`


    // Order - The order of rules in learning module rule
    Order int `json:"order"`

}

// String returns a JSON representation of the model
func (o *Learningmoduleruleparts) String() string {
    
    
     o.Value = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmoduleruleparts) MarshalJSON() ([]byte, error) {
    type Alias Learningmoduleruleparts

    if LearningmodulerulepartsMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulerulepartsMarshalled = true

    return json.Marshal(&struct {
        
        Operation string `json:"operation"`
        
        Selector string `json:"selector"`
        
        Value []string `json:"value"`
        
        Order int `json:"order"`
        *Alias
    }{

        


        


        
        Value: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}


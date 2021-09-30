package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnsweroptionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnsweroptionDud struct { 
    


    


    

}

// Answeroption
type Answeroption struct { 
    // Id
    Id string `json:"id"`


    // Text
    Text string `json:"text"`


    // Value
    Value int `json:"value"`

}

// String returns a JSON representation of the model
func (o *Answeroption) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Answeroption) MarshalJSON() ([]byte, error) {
    type Alias Answeroption

    if AnsweroptionMarshalled {
        return []byte("{}"), nil
    }
    AnsweroptionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        Value int `json:"value"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


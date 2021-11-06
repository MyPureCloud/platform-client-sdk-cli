package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SentimentfeedbackentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SentimentfeedbackentitylistingDud struct { 
    

}

// Sentimentfeedbackentitylisting
type Sentimentfeedbackentitylisting struct { 
    // Entities
    Entities []Sentimentfeedback `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Sentimentfeedbackentitylisting) String() string {
    
    
     o.Entities = []Sentimentfeedback{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sentimentfeedbackentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Sentimentfeedbackentitylisting

    if SentimentfeedbackentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SentimentfeedbackentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Sentimentfeedback `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Sentimentfeedback{{}},
        

        
        Alias: (*Alias)(u),
    })
}


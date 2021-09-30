package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NlufeedbackresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NlufeedbackresponseDud struct { 
    Id string `json:"id"`


    


    


    Version Nludomainversion `json:"version"`


    DateCreated time.Time `json:"dateCreated"`


    SelfUri string `json:"selfUri"`

}

// Nlufeedbackresponse
type Nlufeedbackresponse struct { 
    


    // Text - The feedback text.
    Text string `json:"text"`


    // Intents - Detected intent of the utterance
    Intents []Intentfeedback `json:"intents"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Nlufeedbackresponse) String() string {
    
    
    
    
    
    
    
    
     o.Intents = []Intentfeedback{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nlufeedbackresponse) MarshalJSON() ([]byte, error) {
    type Alias Nlufeedbackresponse

    if NlufeedbackresponseMarshalled {
        return []byte("{}"), nil
    }
    NlufeedbackresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Text string `json:"text"`
        
        Intents []Intentfeedback `json:"intents"`
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Intents: []Intentfeedback{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


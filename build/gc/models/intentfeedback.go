package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntentfeedbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntentfeedbackDud struct { 
    Name string `json:"name"`


    Probability float64 `json:"probability"`


    Entities []Detectednamedentity `json:"entities"`


    

}

// Intentfeedback
type Intentfeedback struct { 
    


    


    


    // Assessment - The assessment on the detection for feedback text.
    Assessment string `json:"assessment"`

}

// String returns a JSON representation of the model
func (o *Intentfeedback) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intentfeedback) MarshalJSON() ([]byte, error) {
    type Alias Intentfeedback

    if IntentfeedbackMarshalled {
        return []byte("{}"), nil
    }
    IntentfeedbackMarshalled = true

    return json.Marshal(&struct {
        
        Assessment string `json:"assessment"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


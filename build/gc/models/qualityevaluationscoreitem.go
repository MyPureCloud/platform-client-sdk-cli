package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QualityevaluationscoreitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QualityevaluationscoreitemDud struct { 
    EvaluationId string `json:"evaluationId"`


    ConversationId string `json:"conversationId"`


    ConversationDate time.Time `json:"conversationDate"`


    ConversationEndDate time.Time `json:"conversationEndDate"`


    FormName string `json:"formName"`


    Points int `json:"points"`


    EvaluationScore float64 `json:"evaluationScore"`


    MaxPoints int `json:"maxPoints"`


    

}

// Qualityevaluationscoreitem
type Qualityevaluationscoreitem struct { 
    


    


    


    


    


    


    


    


    // MediaTypes - A list of media types for the metric
    MediaTypes []string `json:"mediaTypes"`

}

// String returns a JSON representation of the model
func (o *Qualityevaluationscoreitem) String() string {
     o.MediaTypes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Qualityevaluationscoreitem) MarshalJSON() ([]byte, error) {
    type Alias Qualityevaluationscoreitem

    if QualityevaluationscoreitemMarshalled {
        return []byte("{}"), nil
    }
    QualityevaluationscoreitemMarshalled = true

    return json.Marshal(&struct {
        
        MediaTypes []string `json:"mediaTypes"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        MediaTypes: []string{""},
        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeanswerdocumentsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeanswerdocumentsresponseDud struct { 
    


    

}

// Knowledgeanswerdocumentsresponse
type Knowledgeanswerdocumentsresponse struct { 
    // Results - The results with answers if the answerMode request property is not set or contains \"AnswerHighlight\". Empty array otherwise.
    Results []Knowledgeanswerdocumentresponse `json:"results"`


    // AnswerGeneration - The results with AI-generated answer if the answerMode request property contains \"AnswerGeneration\".
    AnswerGeneration Knowledgeanswergenerationresponse `json:"answerGeneration"`

}

// String returns a JSON representation of the model
func (o *Knowledgeanswerdocumentsresponse) String() string {
     o.Results = []Knowledgeanswerdocumentresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeanswerdocumentsresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeanswerdocumentsresponse

    if KnowledgeanswerdocumentsresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeanswerdocumentsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Knowledgeanswerdocumentresponse `json:"results"`
        
        AnswerGeneration Knowledgeanswergenerationresponse `json:"answerGeneration"`
        *Alias
    }{

        
        Results: []Knowledgeanswerdocumentresponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}


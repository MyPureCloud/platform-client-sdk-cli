package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgegenerationsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgegenerationsettingDud struct { 
    


    

}

// Knowledgegenerationsetting
type Knowledgegenerationsetting struct { 
    // AnswerGeneration - Indicates if answer generation is enabled for the setting.
    AnswerGeneration bool `json:"answerGeneration"`


    // GenerationLanguage - Answer generation language.
    GenerationLanguage string `json:"generationLanguage"`

}

// String returns a JSON representation of the model
func (o *Knowledgegenerationsetting) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgegenerationsetting) MarshalJSON() ([]byte, error) {
    type Alias Knowledgegenerationsetting

    if KnowledgegenerationsettingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgegenerationsettingMarshalled = true

    return json.Marshal(&struct {
        
        AnswerGeneration bool `json:"answerGeneration"`
        
        GenerationLanguage string `json:"generationLanguage"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManualsearchconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManualsearchconfigDud struct { 
    


    

}

// Manualsearchconfig
type Manualsearchconfig struct { 
    // ArticlesWithAnswerHighlights - Articles with answer highlights.
    ArticlesWithAnswerHighlights bool `json:"articlesWithAnswerHighlights"`


    // AnswerGeneration - Answer generation.
    AnswerGeneration bool `json:"answerGeneration"`

}

// String returns a JSON representation of the model
func (o *Manualsearchconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Manualsearchconfig) MarshalJSON() ([]byte, error) {
    type Alias Manualsearchconfig

    if ManualsearchconfigMarshalled {
        return []byte("{}"), nil
    }
    ManualsearchconfigMarshalled = true

    return json.Marshal(&struct {
        
        ArticlesWithAnswerHighlights bool `json:"articlesWithAnswerHighlights"`
        
        AnswerGeneration bool `json:"answerGeneration"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


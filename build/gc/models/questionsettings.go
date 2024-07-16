package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuestionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuestionsettingsDud struct { 
    


    

}

// Questionsettings
type Questionsettings struct { 
    // QuestionIndex - This field represents the location of the Question in the form. Note: Indexes are zero-based
    QuestionIndex int `json:"questionIndex"`


    // Settings
    Settings Aiscoringsetting `json:"settings"`

}

// String returns a JSON representation of the model
func (o *Questionsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Questionsettings) MarshalJSON() ([]byte, error) {
    type Alias Questionsettings

    if QuestionsettingsMarshalled {
        return []byte("{}"), nil
    }
    QuestionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        QuestionIndex int `json:"questionIndex"`
        
        Settings Aiscoringsetting `json:"settings"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


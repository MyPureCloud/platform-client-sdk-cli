package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuestiongroupsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuestiongroupsettingsDud struct { 
    


    

}

// Questiongroupsettings
type Questiongroupsettings struct { 
    // QuestionGroupIndex - This field represents the location of the Question Group in the form. Note: Indexes are zero-based
    QuestionGroupIndex int `json:"questionGroupIndex"`


    // QuestionSettings
    QuestionSettings []Questionsettings `json:"questionSettings"`

}

// String returns a JSON representation of the model
func (o *Questiongroupsettings) String() string {
    
     o.QuestionSettings = []Questionsettings{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Questiongroupsettings) MarshalJSON() ([]byte, error) {
    type Alias Questiongroupsettings

    if QuestiongroupsettingsMarshalled {
        return []byte("{}"), nil
    }
    QuestiongroupsettingsMarshalled = true

    return json.Marshal(&struct {
        
        QuestionGroupIndex int `json:"questionGroupIndex"`
        
        QuestionSettings []Questionsettings `json:"questionSettings"`
        *Alias
    }{

        


        
        QuestionSettings: []Questionsettings{{}},
        

        Alias: (*Alias)(u),
    })
}


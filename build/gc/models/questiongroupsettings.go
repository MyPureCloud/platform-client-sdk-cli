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
    // QuestionGroupContextId - The context id of the question group in the form.
    QuestionGroupContextId string `json:"questionGroupContextId"`


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
        
        QuestionGroupContextId string `json:"questionGroupContextId"`
        
        QuestionSettings []Questionsettings `json:"questionSettings"`
        *Alias
    }{

        


        
        QuestionSettings: []Questionsettings{{}},
        

        Alias: (*Alias)(u),
    })
}


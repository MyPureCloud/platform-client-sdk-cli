package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictivescoringsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictivescoringsettingsDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Predictivescoringsettings
type Predictivescoringsettings struct { 
    


    // QuestionGroupSettings
    QuestionGroupSettings []Questiongroupsettings `json:"questionGroupSettings"`


    

}

// String returns a JSON representation of the model
func (o *Predictivescoringsettings) String() string {
     o.QuestionGroupSettings = []Questiongroupsettings{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictivescoringsettings) MarshalJSON() ([]byte, error) {
    type Alias Predictivescoringsettings

    if PredictivescoringsettingsMarshalled {
        return []byte("{}"), nil
    }
    PredictivescoringsettingsMarshalled = true

    return json.Marshal(&struct {
        
        QuestionGroupSettings []Questiongroupsettings `json:"questionGroupSettings"`
        *Alias
    }{

        


        
        QuestionGroupSettings: []Questiongroupsettings{{}},
        


        

        Alias: (*Alias)(u),
    })
}


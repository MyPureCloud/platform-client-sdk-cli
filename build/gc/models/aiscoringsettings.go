package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AiscoringsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AiscoringsettingsDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Aiscoringsettings
type Aiscoringsettings struct { 
    


    // QuestionGroupSettings
    QuestionGroupSettings []Questiongroupsettings `json:"questionGroupSettings"`


    

}

// String returns a JSON representation of the model
func (o *Aiscoringsettings) String() string {
     o.QuestionGroupSettings = []Questiongroupsettings{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aiscoringsettings) MarshalJSON() ([]byte, error) {
    type Alias Aiscoringsettings

    if AiscoringsettingsMarshalled {
        return []byte("{}"), nil
    }
    AiscoringsettingsMarshalled = true

    return json.Marshal(&struct {
        
        QuestionGroupSettings []Questiongroupsettings `json:"questionGroupSettings"`
        *Alias
    }{

        


        
        QuestionGroupSettings: []Questiongroupsettings{{}},
        


        

        Alias: (*Alias)(u),
    })
}


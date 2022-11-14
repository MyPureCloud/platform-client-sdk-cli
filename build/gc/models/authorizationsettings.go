package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthorizationsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthorizationsettingsDud struct { 
    Id string `json:"id"`


    AnalysisEnabled bool `json:"analysisEnabled"`


    AnalysisDays int `json:"analysisDays"`


    SelfUri string `json:"selfUri"`

}

// Authorizationsettings
type Authorizationsettings struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Authorizationsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authorizationsettings) MarshalJSON() ([]byte, error) {
    type Alias Authorizationsettings

    if AuthorizationsettingsMarshalled {
        return []byte("{}"), nil
    }
    AuthorizationsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


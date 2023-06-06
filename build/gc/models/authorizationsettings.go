package models
import (
    "time"
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


    


    


    DateLastCalculated time.Time `json:"dateLastCalculated"`


    DateLastActive time.Time `json:"dateLastActive"`


    SelfUri string `json:"selfUri"`

}

// Authorizationsettings
type Authorizationsettings struct { 
    


    // AnalysisEnabled - Boolean showing if organization is opted in or not to unused role/perm analysis
    AnalysisEnabled bool `json:"analysisEnabled"`


    // AnalysisDays - Integer number of days to analyze user usage
    AnalysisDays int `json:"analysisDays"`


    


    


    

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
        
        AnalysisEnabled bool `json:"analysisEnabled"`
        
        AnalysisDays int `json:"analysisDays"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


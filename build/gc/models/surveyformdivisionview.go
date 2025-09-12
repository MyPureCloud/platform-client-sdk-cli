package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyformdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyformdivisionviewDud struct { 
    Id string `json:"id"`


    


    


    ContextId string `json:"contextId"`


    SelfUri string `json:"selfUri"`

}

// Surveyformdivisionview
type Surveyformdivisionview struct { 
    


    // Name
    Name string `json:"name"`


    // Language - Language for survey viewer localization. Currently localized languages: da, de, en-US, es, fi, fr, it, ja, ko, nl, no, pl, pt-BR, sv, th, tr, zh-CH, zh-TW
    Language string `json:"language"`


    


    

}

// String returns a JSON representation of the model
func (o *Surveyformdivisionview) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyformdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Surveyformdivisionview

    if SurveyformdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    SurveyformdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Language string `json:"language"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


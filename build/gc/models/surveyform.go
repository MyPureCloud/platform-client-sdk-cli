package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyformMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyformDud struct { 
    Id string `json:"id"`


    


    ModifiedDate time.Time `json:"modifiedDate"`


    Published bool `json:"published"`


    


    ContextId string `json:"contextId"`


    


    


    


    


    PublishedVersions Domainentitylistingsurveyform `json:"publishedVersions"`


    SelfUri string `json:"selfUri"`

}

// Surveyform
type Surveyform struct { 
    


    // Name - The survey form name
    Name string `json:"name"`


    


    


    // Disabled - Is this form disabled
    Disabled bool `json:"disabled"`


    


    // Language - Language for survey viewer localization. Currently localized languages: da, de, en-US, es, fi, fr, it, ja, ko, nl, no, pl, pt-BR, sv, th, tr, zh-CH, zh-TW
    Language string `json:"language"`


    // Header - Markdown text for the top of the form.
    Header string `json:"header"`


    // Footer - Markdown text for the bottom of the form.
    Footer string `json:"footer"`


    // QuestionGroups - A list of question groups
    QuestionGroups []Surveyquestiongroup `json:"questionGroups"`


    


    

}

// String returns a JSON representation of the model
func (o *Surveyform) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.QuestionGroups = []Surveyquestiongroup{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyform) MarshalJSON() ([]byte, error) {
    type Alias Surveyform

    if SurveyformMarshalled {
        return []byte("{}"), nil
    }
    SurveyformMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        Disabled bool `json:"disabled"`
        
        
        
        Language string `json:"language"`
        
        Header string `json:"header"`
        
        Footer string `json:"footer"`
        
        QuestionGroups []Surveyquestiongroup `json:"questionGroups"`
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        QuestionGroups: []Surveyquestiongroup{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


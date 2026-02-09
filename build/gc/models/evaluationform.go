package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationformMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationformDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    AiScoring Aiscoringsettings `json:"aiScoring"`


    


    SelfUri string `json:"selfUri"`

}

// Evaluationform
type Evaluationform struct { 
    


    // Name - The evaluation form name
    Name string `json:"name"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // Published
    Published bool `json:"published"`


    // ContextId
    ContextId string `json:"contextId"`


    // QuestionGroups - A list of question groups
    QuestionGroups []Evaluationquestiongroup `json:"questionGroups"`


    // PublishedVersions - A list of the published versions of this form. Not populated by default, its availability depends on the endpoint. Use the 'expand=publishHistory' query parameter to retrieve this data where applicable (refer to the endpoint description to see if it is applicable).
    PublishedVersions Domainentitylistingevaluationform `json:"publishedVersions"`


    // EvaluationSettings - Settings for evaluations associated with this form
    EvaluationSettings Evaluationsettings `json:"evaluationSettings"`


    // LatestVersionFormName - The name of the form's most recently published version
    LatestVersionFormName string `json:"latestVersionFormName"`


    


    // Dialect - The language dialect for this evaluation form. Supported dialects: ar, cs, da, de, en-US, es, fi, fr, fr-CA, he, hi, it, ja, ko, nl, no, pl, pt-BR, pt-PT, ru, sv, th, tr, uk, zh-CN, zh-TW
    Dialect string `json:"dialect"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationform) String() string {
    
    
    
    
     o.QuestionGroups = []Evaluationquestiongroup{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationform) MarshalJSON() ([]byte, error) {
    type Alias Evaluationform

    if EvaluationformMarshalled {
        return []byte("{}"), nil
    }
    EvaluationformMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        Published bool `json:"published"`
        
        ContextId string `json:"contextId"`
        
        QuestionGroups []Evaluationquestiongroup `json:"questionGroups"`
        
        PublishedVersions Domainentitylistingevaluationform `json:"publishedVersions"`
        
        EvaluationSettings Evaluationsettings `json:"evaluationSettings"`
        
        LatestVersionFormName string `json:"latestVersionFormName"`
        
        Dialect string `json:"dialect"`
        *Alias
    }{

        


        


        


        


        


        
        QuestionGroups: []Evaluationquestiongroup{{}},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


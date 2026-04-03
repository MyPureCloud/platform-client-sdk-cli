package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationformsearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationformsearchresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    AiScoring Aiscoringsettings `json:"aiScoring"`


    


    SelfUri string `json:"selfUri"`

}

// Evaluationformsearchresponse
type Evaluationformsearchresponse struct { 
    


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // Published
    Published bool `json:"published"`


    // ContextId
    ContextId string `json:"contextId"`


    // QuestionGroups - A list of question groups
    QuestionGroups []Evaluationquestiongroup `json:"questionGroups"`


    // WeightMode - Mode for evaluation form weight
    WeightMode string `json:"weightMode"`


    // EvaluationSettings - Settings for evaluations associated with this form
    EvaluationSettings Evaluationsettings `json:"evaluationSettings"`


    // PublishedVersions - A list of the published versions of this form. Not populated by default, its availability depends on the endpoint. Use the 'expand=publishHistory' query parameter to retrieve this data where applicable (refer to the endpoint description to see if it is applicable).
    PublishedVersions Domainentitylistingevaluationform `json:"publishedVersions"`


    // LatestVersionFormName - The name of the form's most recently published version
    LatestVersionFormName string `json:"latestVersionFormName"`


    


    // Dialect - The language dialect for this evaluation form. Supported dialects: ar, cs, da, de, en-US, es, fi, fr, fr-CA, he, hi, it, ja, ko, nl, no, pl, pt-BR, pt-PT, ru, sv, th, tr, uk, zh-CN, zh-TW
    Dialect string `json:"dialect"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationformsearchresponse) String() string {
    
    
    
     o.QuestionGroups = []Evaluationquestiongroup{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationformsearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Evaluationformsearchresponse

    if EvaluationformsearchresponseMarshalled {
        return []byte("{}"), nil
    }
    EvaluationformsearchresponseMarshalled = true

    return json.Marshal(&struct {
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        Published bool `json:"published"`
        
        ContextId string `json:"contextId"`
        
        QuestionGroups []Evaluationquestiongroup `json:"questionGroups"`
        
        WeightMode string `json:"weightMode"`
        
        EvaluationSettings Evaluationsettings `json:"evaluationSettings"`
        
        PublishedVersions Domainentitylistingevaluationform `json:"publishedVersions"`
        
        LatestVersionFormName string `json:"latestVersionFormName"`
        
        Dialect string `json:"dialect"`
        *Alias
    }{

        


        


        


        


        
        QuestionGroups: []Evaluationquestiongroup{{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HealthinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HealthinfoDud struct { 
    


    


    


    


    


    

}

// Healthinfo
type Healthinfo struct { 
    // Status - Status of health computation for this intent.
    Status string `json:"status"`


    // ErrorInfo - Error details for the intent, if any.
    ErrorInfo Flowhealtherrorinfo `json:"errorInfo"`


    // OverallScore - Overall health score for the intent ranged between 0 and 100 as 100 is the perfect health score.
    OverallScore float32 `json:"overallScore"`


    // IssueCount - Number of issues found in the intent.
    IssueCount int `json:"issueCount"`


    // StaticValidationResults - Validation results for the intent.
    StaticValidationResults []string `json:"staticValidationResults"`


    // Utterances - Utterances for this intent.
    Utterances []Flowhealthintentutterance `json:"utterances"`

}

// String returns a JSON representation of the model
func (o *Healthinfo) String() string {
    
    
    
    
     o.StaticValidationResults = []string{""} 
     o.Utterances = []Flowhealthintentutterance{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Healthinfo) MarshalJSON() ([]byte, error) {
    type Alias Healthinfo

    if HealthinfoMarshalled {
        return []byte("{}"), nil
    }
    HealthinfoMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        ErrorInfo Flowhealtherrorinfo `json:"errorInfo"`
        
        OverallScore float32 `json:"overallScore"`
        
        IssueCount int `json:"issueCount"`
        
        StaticValidationResults []string `json:"staticValidationResults"`
        
        Utterances []Flowhealthintentutterance `json:"utterances"`
        *Alias
    }{

        


        


        


        


        
        StaticValidationResults: []string{""},
        


        
        Utterances: []Flowhealthintentutterance{{}},
        

        Alias: (*Alias)(u),
    })
}


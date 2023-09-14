package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocalehealthMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocalehealthDud struct { 
    


    


    

}

// Localehealth
type Localehealth struct { 
    // OverallScore - Overall health score for the intent.
    OverallScore float32 `json:"overallScore"`


    // IssueCount - Number of issues found in the intent.
    IssueCount int `json:"issueCount"`


    // StaticValidationResults - Validation results for the intent.
    StaticValidationResults []string `json:"staticValidationResults"`

}

// String returns a JSON representation of the model
func (o *Localehealth) String() string {
    
    
     o.StaticValidationResults = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Localehealth) MarshalJSON() ([]byte, error) {
    type Alias Localehealth

    if LocalehealthMarshalled {
        return []byte("{}"), nil
    }
    LocalehealthMarshalled = true

    return json.Marshal(&struct {
        
        OverallScore float32 `json:"overallScore"`
        
        IssueCount int `json:"issueCount"`
        
        StaticValidationResults []string `json:"staticValidationResults"`
        *Alias
    }{

        


        


        
        StaticValidationResults: []string{""},
        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowhealthutteranceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowhealthutteranceDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowhealthutterance
type Flowhealthutterance struct { 
    


    // Text - Utterance Text.
    Text string `json:"text"`


    // IssueCount - Number of issues found for this utterance.
    IssueCount int `json:"issueCount"`


    // Language - Language provided for this utterance's health.
    Language string `json:"language"`


    // StaticValidationResults - Validation results for the utterance.
    StaticValidationResults []string `json:"staticValidationResults"`


    // OutlierInfo - Details about this utterance being an outlier or not.
    OutlierInfo Outlierinfo `json:"outlierInfo"`


    // ConfusionInfo - Confusion details with other utterances.
    ConfusionInfo Confusiondetails `json:"confusionInfo"`


    

}

// String returns a JSON representation of the model
func (o *Flowhealthutterance) String() string {
    
    
    
     o.StaticValidationResults = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowhealthutterance) MarshalJSON() ([]byte, error) {
    type Alias Flowhealthutterance

    if FlowhealthutteranceMarshalled {
        return []byte("{}"), nil
    }
    FlowhealthutteranceMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        IssueCount int `json:"issueCount"`
        
        Language string `json:"language"`
        
        StaticValidationResults []string `json:"staticValidationResults"`
        
        OutlierInfo Outlierinfo `json:"outlierInfo"`
        
        ConfusionInfo Confusiondetails `json:"confusionInfo"`
        *Alias
    }{

        


        


        


        


        
        StaticValidationResults: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


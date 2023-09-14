package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowhealthintentutteranceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowhealthintentutteranceDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Flowhealthintentutterance
type Flowhealthintentutterance struct { 
    


    // Text - Utterance Text.
    Text string `json:"text"`


    // IssueCount - Number of issues found for this utterance.
    IssueCount int `json:"issueCount"`


    // StaticValidationResults - Validation results for this utterance.
    StaticValidationResults []string `json:"staticValidationResults"`


    // OutlierInfo - Details about this utterance being an outlier or not.
    OutlierInfo Outlierinfo `json:"outlierInfo"`


    // ConfusionInfo - Confusion details with other utterances.
    ConfusionInfo Confusioninfo `json:"confusionInfo"`

}

// String returns a JSON representation of the model
func (o *Flowhealthintentutterance) String() string {
    
    
     o.StaticValidationResults = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowhealthintentutterance) MarshalJSON() ([]byte, error) {
    type Alias Flowhealthintentutterance

    if FlowhealthintentutteranceMarshalled {
        return []byte("{}"), nil
    }
    FlowhealthintentutteranceMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        IssueCount int `json:"issueCount"`
        
        StaticValidationResults []string `json:"staticValidationResults"`
        
        OutlierInfo Outlierinfo `json:"outlierInfo"`
        
        ConfusionInfo Confusioninfo `json:"confusionInfo"`
        *Alias
    }{

        


        


        


        
        StaticValidationResults: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyerrordetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyerrordetailsDud struct { 
    


    

}

// Surveyerrordetails
type Surveyerrordetails struct { 
    // FlowDiagnosticInfo - Additional information about any errors that occurred in the survey invite flow.
    FlowDiagnosticInfo Flowdiagnosticinfo `json:"flowDiagnosticInfo"`


    // SurveyErrorReason - An error code indicating the reason for the survey failure.
    SurveyErrorReason string `json:"surveyErrorReason"`

}

// String returns a JSON representation of the model
func (o *Surveyerrordetails) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyerrordetails) MarshalJSON() ([]byte, error) {
    type Alias Surveyerrordetails

    if SurveyerrordetailsMarshalled {
        return []byte("{}"), nil
    }
    SurveyerrordetailsMarshalled = true

    return json.Marshal(&struct { 
        FlowDiagnosticInfo Flowdiagnosticinfo `json:"flowDiagnosticInfo"`
        
        SurveyErrorReason string `json:"surveyErrorReason"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


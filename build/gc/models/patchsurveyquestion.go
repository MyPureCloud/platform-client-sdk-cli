package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchsurveyquestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchsurveyquestionDud struct { 
    


    


    


    


    

}

// Patchsurveyquestion
type Patchsurveyquestion struct { 
    // VarType - Type of survey question.
    VarType string `json:"type"`


    // Label - Label of question.
    Label string `json:"label"`


    // CustomerProperty - The customer property that the answer maps to.
    CustomerProperty string `json:"customerProperty"`


    // Choices - Choices available to user.
    Choices []string `json:"choices"`


    // IsMandatory - Whether answering this question is mandatory.
    IsMandatory bool `json:"isMandatory"`

}

// String returns a JSON representation of the model
func (o *Patchsurveyquestion) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Choices = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchsurveyquestion) MarshalJSON() ([]byte, error) {
    type Alias Patchsurveyquestion

    if PatchsurveyquestionMarshalled {
        return []byte("{}"), nil
    }
    PatchsurveyquestionMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Label string `json:"label"`
        
        CustomerProperty string `json:"customerProperty"`
        
        Choices []string `json:"choices"`
        
        IsMandatory bool `json:"isMandatory"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Choices: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}


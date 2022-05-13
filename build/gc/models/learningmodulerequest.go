package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulerequestDud struct { 
    


    


    


    


    


    


    

}

// Learningmodulerequest - Learning module request
type Learningmodulerequest struct { 
    // Name - The name of learning module
    Name string `json:"name"`


    // Description - The description of learning module
    Description string `json:"description"`


    // CompletionTimeInDays - The completion time of learning module in days
    CompletionTimeInDays int `json:"completionTimeInDays"`


    // InformSteps - The list of inform steps in a learning module
    InformSteps []Learningmoduleinformsteprequest `json:"informSteps"`


    // VarType - The type for the learning module
    VarType string `json:"type"`


    // AssessmentForm - The assessment form for learning module
    AssessmentForm Assessmentform `json:"assessmentForm"`


    // CoverArt - The cover art for the learning module
    CoverArt Learningmodulecoverartrequest `json:"coverArt"`

}

// String returns a JSON representation of the model
func (o *Learningmodulerequest) String() string {
    
    
    
     o.InformSteps = []Learningmoduleinformsteprequest{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulerequest) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulerequest

    if LearningmodulerequestMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        CompletionTimeInDays int `json:"completionTimeInDays"`
        
        InformSteps []Learningmoduleinformsteprequest `json:"informSteps"`
        
        VarType string `json:"type"`
        
        AssessmentForm Assessmentform `json:"assessmentForm"`
        
        CoverArt Learningmodulecoverartrequest `json:"coverArt"`
        *Alias
    }{

        


        


        


        
        InformSteps: []Learningmoduleinformsteprequest{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


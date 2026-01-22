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


    // VarType - The type for the learning module. Informational, AssessedContent and Assessment are deprecated
    VarType string `json:"type"`


    // AssessmentForm - The assessment form for learning module
    AssessmentForm Assessmentform `json:"assessmentForm"`


    // CoverArt - The cover art for the learning module
    CoverArt Learningmodulecoverartrequest `json:"coverArt"`


    // LengthInMinutes - The recommended time in minutes to complete the module
    LengthInMinutes int `json:"lengthInMinutes"`


    // ExcludedFromCatalog - If true, learning module is excluded when retrieving modules for manual assignment
    ExcludedFromCatalog bool `json:"excludedFromCatalog"`


    // ExternalId - The external ID of the learning module. Maximum length: 50 characters.
    ExternalId string `json:"externalId"`


    // EnforceContentOrder - If true, learning module content should be viewed one by one in order
    EnforceContentOrder bool `json:"enforceContentOrder"`


    // ReviewAssessmentResults - Allows to view Assessment results in detail
    ReviewAssessmentResults Reviewassessmentresults `json:"reviewAssessmentResults"`


    // AutoAssign - The configuration for linking a module to a rule
    AutoAssign Learningmoduleautoassignrequest `json:"autoAssign"`

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
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        ExcludedFromCatalog bool `json:"excludedFromCatalog"`
        
        ExternalId string `json:"externalId"`
        
        EnforceContentOrder bool `json:"enforceContentOrder"`
        
        ReviewAssessmentResults Reviewassessmentresults `json:"reviewAssessmentResults"`
        
        AutoAssign Learningmoduleautoassignrequest `json:"autoAssign"`
        *Alias
    }{

        


        


        


        
        InformSteps: []Learningmoduleinformsteprequest{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignedlearningmoduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignedlearningmoduleDud struct { 
    Id string `json:"id"`


    


    


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    ModifiedBy Userreference `json:"modifiedBy"`


    DateModified time.Time `json:"dateModified"`


    Version int `json:"version"`


    ExternalId string `json:"externalId"`


    Source string `json:"source"`


    Rule Learningmodulerule `json:"rule"`


    


    SelfUri string `json:"selfUri"`


    IsArchived bool `json:"isArchived"`


    IsPublished bool `json:"isPublished"`


    


    


    


    


    


    


    


    


    


    

}

// Assignedlearningmodule - Learning module response
type Assignedlearningmodule struct { 
    


    // Name - The name of learning module
    Name string `json:"name"`


    // ExcludedFromCatalog - If true, learning module is excluded when retrieving modules for manual assignment
    ExcludedFromCatalog bool `json:"excludedFromCatalog"`


    


    


    


    


    


    


    


    


    // CurrentAssignments - The current assignments for the requested users
    CurrentAssignments []Learningassignment `json:"currentAssignments"`


    


    


    


    // Description - The description of learning module
    Description string `json:"description"`


    // CompletionTimeInDays - The completion time of learning module in days
    CompletionTimeInDays int `json:"completionTimeInDays"`


    // VarType - The type for the learning module
    VarType string `json:"type"`


    // InformSteps - The list of inform steps in a learning module
    InformSteps []Learningmoduleinformstep `json:"informSteps"`


    // AssessmentForm - The assessment form for learning module
    AssessmentForm Assessmentform `json:"assessmentForm"`


    // SummaryData - The learning module summary data
    SummaryData Learningmodulesummary `json:"summaryData"`


    // ReassignSummaryData - The learning module reassign summary data
    ReassignSummaryData Learningmodulereassignsummary `json:"reassignSummaryData"`


    // CoverArt - The cover art for the learning module
    CoverArt Learningmodulecoverartresponse `json:"coverArt"`


    // LengthInMinutes - The recommended time in minutes to complete the module
    LengthInMinutes int `json:"lengthInMinutes"`


    // ArchivalMode - The mode of archival for learning module
    ArchivalMode string `json:"archivalMode"`

}

// String returns a JSON representation of the model
func (o *Assignedlearningmodule) String() string {
    
    
     o.CurrentAssignments = []Learningassignment{{}} 
    
    
    
     o.InformSteps = []Learningmoduleinformstep{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignedlearningmodule) MarshalJSON() ([]byte, error) {
    type Alias Assignedlearningmodule

    if AssignedlearningmoduleMarshalled {
        return []byte("{}"), nil
    }
    AssignedlearningmoduleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ExcludedFromCatalog bool `json:"excludedFromCatalog"`
        
        CurrentAssignments []Learningassignment `json:"currentAssignments"`
        
        Description string `json:"description"`
        
        CompletionTimeInDays int `json:"completionTimeInDays"`
        
        VarType string `json:"type"`
        
        InformSteps []Learningmoduleinformstep `json:"informSteps"`
        
        AssessmentForm Assessmentform `json:"assessmentForm"`
        
        SummaryData Learningmodulesummary `json:"summaryData"`
        
        ReassignSummaryData Learningmodulereassignsummary `json:"reassignSummaryData"`
        
        CoverArt Learningmodulecoverartresponse `json:"coverArt"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        ArchivalMode string `json:"archivalMode"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        CurrentAssignments: []Learningassignment{{}},
        


        


        


        


        


        


        


        
        InformSteps: []Learningmoduleinformstep{{}},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


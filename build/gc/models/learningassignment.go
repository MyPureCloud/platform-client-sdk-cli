package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentDud struct { 
    Id string `json:"id"`


    


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    ModifiedBy Userreference `json:"modifiedBy"`


    DateModified time.Time `json:"dateModified"`


    IsOverdue bool `json:"isOverdue"`


    PercentageScore float32 `json:"percentageScore"`


    IsRule bool `json:"isRule"`


    IsManual bool `json:"isManual"`


    IsPassed bool `json:"isPassed"`


    SelfUri string `json:"selfUri"`


    


    


    


    


    


    

}

// Learningassignment - Learning module assignment with user information
type Learningassignment struct { 
    


    // Assessment - The assessment associated with this assignment
    Assessment Learningassessment `json:"assessment"`


    


    


    


    


    


    


    


    


    


    


    // State - The Learning Assignment state
    State string `json:"state"`


    // DateRecommendedForCompletion - The recommended completion date of the assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateRecommendedForCompletion time.Time `json:"dateRecommendedForCompletion"`


    // Version - The version of Learning module assigned
    Version int `json:"version"`


    // Module - The Learning module object associated with this assignment
    Module Learningmodule `json:"module"`


    // User - The user to whom the assignment is assigned
    User Userreference `json:"user"`


    // AssessmentForm - The assessment form associated with this assignment
    AssessmentForm Assessmentform `json:"assessmentForm"`

}

// String returns a JSON representation of the model
func (o *Learningassignment) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignment) MarshalJSON() ([]byte, error) {
    type Alias Learningassignment

    if LearningassignmentMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentMarshalled = true

    return json.Marshal(&struct {
        
        Assessment Learningassessment `json:"assessment"`
        
        State string `json:"state"`
        
        DateRecommendedForCompletion time.Time `json:"dateRecommendedForCompletion"`
        
        Version int `json:"version"`
        
        Module Learningmodule `json:"module"`
        
        User Userreference `json:"user"`
        
        AssessmentForm Assessmentform `json:"assessmentForm"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


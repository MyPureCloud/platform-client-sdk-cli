package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewgetresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewgetresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Learningmodulepreviewgetresponse - Learning module preview get response
type Learningmodulepreviewgetresponse struct { 
    


    // Name - The name of learning module
    Name string `json:"name"`


    // Description - The description of learning module
    Description string `json:"description"`


    // CoverArt - The cover art for the learning module
    CoverArt Learningmodulecoverartresponse `json:"coverArt"`


    // EnforceContentOrder - If true, learning module content should be viewed one by one in order
    EnforceContentOrder bool `json:"enforceContentOrder"`


    // ReviewAssessmentResults - Allows to view Assessment results in detail
    ReviewAssessmentResults Reviewassessmentresults `json:"reviewAssessmentResults"`


    // AssessmentForm - The assessment form for learning module
    AssessmentForm Assessmentform `json:"assessmentForm"`


    // Assignment - the assignment preview
    Assignment Learningmodulepreviewgetresponseassignment `json:"assignment"`


    

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewgetresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewgetresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewgetresponse

    if LearningmodulepreviewgetresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewgetresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        CoverArt Learningmodulecoverartresponse `json:"coverArt"`
        
        EnforceContentOrder bool `json:"enforceContentOrder"`
        
        ReviewAssessmentResults Reviewassessmentresults `json:"reviewAssessmentResults"`
        
        AssessmentForm Assessmentform `json:"assessmentForm"`
        
        Assignment Learningmodulepreviewgetresponseassignment `json:"assignment"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssessmentjoblistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssessmentjoblistingDud struct { 
    


    


    

}

// Assessmentjoblisting
type Assessmentjoblisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Benefitassessmentjob `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Assessmentjoblisting) String() string {
    
     o.Entities = []Benefitassessmentjob{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assessmentjoblisting) MarshalJSON() ([]byte, error) {
    type Alias Assessmentjoblisting

    if AssessmentjoblistingMarshalled {
        return []byte("{}"), nil
    }
    AssessmentjoblistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Benefitassessmentjob `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Benefitassessmentjob{{}},
        


        

        Alias: (*Alias)(u),
    })
}


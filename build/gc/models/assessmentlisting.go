package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssessmentlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssessmentlistingDud struct { 
    


    


    


    

}

// Assessmentlisting
type Assessmentlisting struct { 
    // Entities
    Entities []Benefitassessment `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Assessmentlisting) String() string {
     o.Entities = []Benefitassessment{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assessmentlisting) MarshalJSON() ([]byte, error) {
    type Alias Assessmentlisting

    if AssessmentlistingMarshalled {
        return []byte("{}"), nil
    }
    AssessmentlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Benefitassessment `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Benefitassessment{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


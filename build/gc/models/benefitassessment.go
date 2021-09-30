package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BenefitassessmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BenefitassessmentDud struct { 
    Id string `json:"id"`


    Queues []Addressableentityref `json:"queues"`


    KpiAssessments []Keyperformanceindicatorassessment `json:"kpiAssessments"`


    State string `json:"state"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Benefitassessment
type Benefitassessment struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Benefitassessment) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Benefitassessment) MarshalJSON() ([]byte, error) {
    type Alias Benefitassessment

    if BenefitassessmentMarshalled {
        return []byte("{}"), nil
    }
    BenefitassessmentMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


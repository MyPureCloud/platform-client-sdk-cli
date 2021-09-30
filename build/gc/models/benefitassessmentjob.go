package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BenefitassessmentjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BenefitassessmentjobDud struct { 
    Id string `json:"id"`


    State string `json:"state"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Benefitassessmentjob
type Benefitassessmentjob struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Benefitassessmentjob) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Benefitassessmentjob) MarshalJSON() ([]byte, error) {
    type Alias Benefitassessmentjob

    if BenefitassessmentjobMarshalled {
        return []byte("{}"), nil
    }
    BenefitassessmentjobMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


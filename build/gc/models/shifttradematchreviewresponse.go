package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradematchreviewresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradematchreviewresponseDud struct { 
    


    


    


    

}

// Shifttradematchreviewresponse
type Shifttradematchreviewresponse struct { 
    // InitiatingUser - Details for the initiatingUser side of the shift trade
    InitiatingUser Shifttradematchreviewuserresponse `json:"initiatingUser"`


    // ReceivingUser - Details for the receivingUser side of the shift trade
    ReceivingUser Shifttradematchreviewuserresponse `json:"receivingUser"`


    // Violations - Constraint violations introduced after being matched that would normally disallow a trade, but which can still be overridden by the shift trade administrator
    Violations []Shifttradematchviolation `json:"violations"`


    // AdminReviewViolations - Constraint violations associated with this shift trade which require shift trade administrator review
    AdminReviewViolations []Shifttradematchviolation `json:"adminReviewViolations"`

}

// String returns a JSON representation of the model
func (o *Shifttradematchreviewresponse) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Violations = []Shifttradematchviolation{{}} 
    
    
    
     o.AdminReviewViolations = []Shifttradematchviolation{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradematchreviewresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradematchreviewresponse

    if ShifttradematchreviewresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradematchreviewresponseMarshalled = true

    return json.Marshal(&struct { 
        InitiatingUser Shifttradematchreviewuserresponse `json:"initiatingUser"`
        
        ReceivingUser Shifttradematchreviewuserresponse `json:"receivingUser"`
        
        Violations []Shifttradematchviolation `json:"violations"`
        
        AdminReviewViolations []Shifttradematchviolation `json:"adminReviewViolations"`
        
        *Alias
    }{
        

        

        

        

        

        
        Violations: []Shifttradematchviolation{{}},
        

        

        
        AdminReviewViolations: []Shifttradematchviolation{{}},
        

        
        Alias: (*Alias)(u),
    })
}


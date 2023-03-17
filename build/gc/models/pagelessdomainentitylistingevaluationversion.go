package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PagelessdomainentitylistingevaluationversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PagelessdomainentitylistingevaluationversionDud struct { 
    


    


    

}

// Pagelessdomainentitylistingevaluationversion
type Pagelessdomainentitylistingevaluationversion struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Evaluationversion `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Pagelessdomainentitylistingevaluationversion) String() string {
    
     o.Entities = []Evaluationversion{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pagelessdomainentitylistingevaluationversion) MarshalJSON() ([]byte, error) {
    type Alias Pagelessdomainentitylistingevaluationversion

    if PagelessdomainentitylistingevaluationversionMarshalled {
        return []byte("{}"), nil
    }
    PagelessdomainentitylistingevaluationversionMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Evaluationversion `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Evaluationversion{{}},
        


        

        Alias: (*Alias)(u),
    })
}


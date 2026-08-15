package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasequeryjobresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasequeryjobresultsresponseDud struct { 
    


    


    


    


    

}

// Casequeryjobresultsresponse
type Casequeryjobresultsresponse struct { 
    // Entities
    Entities []Case `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Casequeryjobresultsresponse) String() string {
     o.Entities = []Case{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casequeryjobresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Casequeryjobresultsresponse

    if CasequeryjobresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    CasequeryjobresultsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Case `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Case{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


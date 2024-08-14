package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsbulkoperationjoblistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsbulkoperationjoblistingDud struct { 
    


    


    

}

// Contactsbulkoperationjoblisting
type Contactsbulkoperationjoblisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Contactsbulkoperationjob `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Contactsbulkoperationjoblisting) String() string {
    
     o.Entities = []Contactsbulkoperationjob{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsbulkoperationjoblisting) MarshalJSON() ([]byte, error) {
    type Alias Contactsbulkoperationjoblisting

    if ContactsbulkoperationjoblistingMarshalled {
        return []byte("{}"), nil
    }
    ContactsbulkoperationjoblistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Contactsbulkoperationjob `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Contactsbulkoperationjob{{}},
        


        

        Alias: (*Alias)(u),
    })
}


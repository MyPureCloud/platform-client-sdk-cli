package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseassociationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseassociationlistingDud struct { 
    


    


    


    

}

// Caseassociationlisting
type Caseassociationlisting struct { 
    // Entities
    Entities []Caseassociation `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Caseassociationlisting) String() string {
     o.Entities = []Caseassociation{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseassociationlisting) MarshalJSON() ([]byte, error) {
    type Alias Caseassociationlisting

    if CaseassociationlistingMarshalled {
        return []byte("{}"), nil
    }
    CaseassociationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Caseassociation `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Caseassociation{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


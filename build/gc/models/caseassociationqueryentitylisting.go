package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseassociationqueryentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseassociationqueryentitylistingDud struct { 
    


    


    


    


    

}

// Caseassociationqueryentitylisting
type Caseassociationqueryentitylisting struct { 
    // Entities
    Entities []Caseassociation `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // After
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Caseassociationqueryentitylisting) String() string {
     o.Entities = []Caseassociation{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseassociationqueryentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Caseassociationqueryentitylisting

    if CaseassociationqueryentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CaseassociationqueryentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Caseassociation `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Caseassociation{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


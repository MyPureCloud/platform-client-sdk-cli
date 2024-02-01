package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperationlistingDud struct { 
    


    


    


    

}

// Operationlisting
type Operationlisting struct { 
    // Entities
    Entities []Operationresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Operationlisting) String() string {
     o.Entities = []Operationresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operationlisting) MarshalJSON() ([]byte, error) {
    type Alias Operationlisting

    if OperationlistingMarshalled {
        return []byte("{}"), nil
    }
    OperationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Operationresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Operationresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


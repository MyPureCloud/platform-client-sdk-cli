package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentvariationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentvariationlistingDud struct { 
    


    


    


    

}

// Documentvariationlisting
type Documentvariationlisting struct { 
    // Entities
    Entities []Documentvariation `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Documentvariationlisting) String() string {
     o.Entities = []Documentvariation{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentvariationlisting) MarshalJSON() ([]byte, error) {
    type Alias Documentvariationlisting

    if DocumentvariationlistingMarshalled {
        return []byte("{}"), nil
    }
    DocumentvariationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Documentvariation `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Documentvariation{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


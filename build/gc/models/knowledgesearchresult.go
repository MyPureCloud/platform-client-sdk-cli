package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesearchresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesearchresultDud struct { 
    


    

}

// Knowledgesearchresult
type Knowledgesearchresult struct { 
    // GeneratedAnswer - The generated answer for search query.
    GeneratedAnswer string `json:"generatedAnswer"`


    // RetrievedReferences - The retrieved references for the search query.
    RetrievedReferences []Knowledgeretrievedreference `json:"retrievedReferences"`

}

// String returns a JSON representation of the model
func (o *Knowledgesearchresult) String() string {
    
     o.RetrievedReferences = []Knowledgeretrievedreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesearchresult) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesearchresult

    if KnowledgesearchresultMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesearchresultMarshalled = true

    return json.Marshal(&struct {
        
        GeneratedAnswer string `json:"generatedAnswer"`
        
        RetrievedReferences []Knowledgeretrievedreference `json:"retrievedReferences"`
        *Alias
    }{

        


        
        RetrievedReferences: []Knowledgeretrievedreference{{}},
        

        Alias: (*Alias)(u),
    })
}


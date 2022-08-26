package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgecontextvaluereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgecontextvaluereferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Knowledgecontextvaluereference
type Knowledgecontextvaluereference struct { 
    // Id - The globally unique identifier for the knowledge context value.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgecontextvaluereference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgecontextvaluereference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgecontextvaluereference

    if KnowledgecontextvaluereferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgecontextvaluereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgecontextreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgecontextreferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Knowledgecontextreference
type Knowledgecontextreference struct { 
    // Id - The globally unique identifier for the knowledge context.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgecontextreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgecontextreference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgecontextreference

    if KnowledgecontextreferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgecontextreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


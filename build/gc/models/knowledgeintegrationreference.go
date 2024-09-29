package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeintegrationreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeintegrationreferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Knowledgeintegrationreference
type Knowledgeintegrationreference struct { 
    // Id - The globally unique identifier for the integration.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgeintegrationreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeintegrationreference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeintegrationreference

    if KnowledgeintegrationreferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeintegrationreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeoperationsourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeoperationsourceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Knowledgeoperationsource
type Knowledgeoperationsource struct { 
    


    // VarType - The source type.
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgeoperationsource) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeoperationsource) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeoperationsource

    if KnowledgeoperationsourceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeoperationsourceMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


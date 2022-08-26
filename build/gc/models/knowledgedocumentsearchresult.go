package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentsearchresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentsearchresultDud struct { 
    


    

}

// Knowledgedocumentsearchresult
type Knowledgedocumentsearchresult struct { 
    // Confidence - The confidence associated with a document with respect to a search query.
    Confidence float64 `json:"confidence"`


    // Document - Document that matched the query.
    Document Knowledgedocumentresponse `json:"document"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsearchresult) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentsearchresult) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentsearchresult

    if KnowledgedocumentsearchresultMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentsearchresultMarshalled = true

    return json.Marshal(&struct {
        
        Confidence float64 `json:"confidence"`
        
        Document Knowledgedocumentresponse `json:"document"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentguestsearchresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentguestsearchresultDud struct { 
    


    

}

// Knowledgedocumentguestsearchresult
type Knowledgedocumentguestsearchresult struct { 
    // Confidence - The confidence associated with a document with respect to a search query.
    Confidence float64 `json:"confidence"`


    // Document - Document that matched the query.
    Document Knowledgeguestdocument `json:"document"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentguestsearchresult) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentguestsearchresult) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentguestsearchresult

    if KnowledgedocumentguestsearchresultMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentguestsearchresultMarshalled = true

    return json.Marshal(&struct {
        
        Confidence float64 `json:"confidence"`
        
        Document Knowledgeguestdocument `json:"document"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


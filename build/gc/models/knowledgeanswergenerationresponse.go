package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeanswergenerationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeanswergenerationresponseDud struct { 
    


    

}

// Knowledgeanswergenerationresponse
type Knowledgeanswergenerationresponse struct { 
    // Answer - The AI-generated answer.
    Answer string `json:"answer"`


    // Documents - The documents used for answer generation.
    Documents []Answergenerationdocument `json:"documents"`

}

// String returns a JSON representation of the model
func (o *Knowledgeanswergenerationresponse) String() string {
    
     o.Documents = []Answergenerationdocument{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeanswergenerationresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeanswergenerationresponse

    if KnowledgeanswergenerationresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeanswergenerationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Answer string `json:"answer"`
        
        Documents []Answergenerationdocument `json:"documents"`
        *Alias
    }{

        


        
        Documents: []Answergenerationdocument{{}},
        

        Alias: (*Alias)(u),
    })
}


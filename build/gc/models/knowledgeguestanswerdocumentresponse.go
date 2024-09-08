package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestanswerdocumentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestanswerdocumentresponseDud struct { 
    


    


    


    

}

// Knowledgeguestanswerdocumentresponse
type Knowledgeguestanswerdocumentresponse struct { 
    // Id - The document id.
    Id string `json:"id"`


    // Title - The document title.
    Title string `json:"title"`


    // Answer - The answer found inside a variationContent.
    Answer string `json:"answer"`


    // Variation - The variation with the answer's highlight data.
    Variation Knowledgeguestdocumentvariationanswer `json:"variation"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestanswerdocumentresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestanswerdocumentresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestanswerdocumentresponse

    if KnowledgeguestanswerdocumentresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestanswerdocumentresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Answer string `json:"answer"`
        
        Variation Knowledgeguestdocumentvariationanswer `json:"variation"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


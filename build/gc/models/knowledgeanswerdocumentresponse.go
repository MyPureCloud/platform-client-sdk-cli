package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeanswerdocumentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeanswerdocumentresponseDud struct { 
    


    


    


    

}

// Knowledgeanswerdocumentresponse
type Knowledgeanswerdocumentresponse struct { 
    // Id - The document id.
    Id string `json:"id"`


    // Title - The document title.
    Title string `json:"title"`


    // Answer - The answer found inside a variationContent.
    Answer string `json:"answer"`


    // Variation - The variation with the answer's highlight data.
    Variation Documentvariationanswer `json:"variation"`

}

// String returns a JSON representation of the model
func (o *Knowledgeanswerdocumentresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeanswerdocumentresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeanswerdocumentresponse

    if KnowledgeanswerdocumentresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeanswerdocumentresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Answer string `json:"answer"`
        
        Variation Documentvariationanswer `json:"variation"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


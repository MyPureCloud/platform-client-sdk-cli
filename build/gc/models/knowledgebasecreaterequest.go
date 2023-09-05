package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgebasecreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgebasecreaterequestDud struct { 
    


    


    

}

// Knowledgebasecreaterequest
type Knowledgebasecreaterequest struct { 
    // Name - Knowledge base name
    Name string `json:"name"`


    // Description - Knowledge base description
    Description string `json:"description"`


    // CoreLanguage - Core language for knowledge base in which initial content must be created, language codes [en-US, en-UK, en-AU, de-DE] are supported currently, however the new DX knowledge will support all these language codes
    CoreLanguage string `json:"coreLanguage"`

}

// String returns a JSON representation of the model
func (o *Knowledgebasecreaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgebasecreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgebasecreaterequest

    if KnowledgebasecreaterequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgebasecreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        CoreLanguage string `json:"coreLanguage"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


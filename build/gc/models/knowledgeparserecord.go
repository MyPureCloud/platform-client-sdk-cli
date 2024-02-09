package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeparserecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeparserecordDud struct { 
    


    


    

}

// Knowledgeparserecord
type Knowledgeparserecord struct { 
    // Id - Unique id for the parsed data.
    Id string `json:"id"`


    // Title - Parsed article title.
    Title string `json:"title"`


    // Body - Parsed article content.
    Body Documentbody `json:"body"`

}

// String returns a JSON representation of the model
func (o *Knowledgeparserecord) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeparserecord) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeparserecord

    if KnowledgeparserecordMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeparserecordMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Body Documentbody `json:"body"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


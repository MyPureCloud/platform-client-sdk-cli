package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentuploadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentuploadDud struct { 
    


    


    


    

}

// Documentupload
type Documentupload struct { 
    // Name - The name of the document
    Name string `json:"name"`


    // Workspace - The workspace the document will be uploaded to
    Workspace Domainentityref `json:"workspace"`


    // Tags
    Tags []string `json:"tags"`


    // TagIds
    TagIds []string `json:"tagIds"`

}

// String returns a JSON representation of the model
func (o *Documentupload) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Tags = []string{""} 
    
    
    
     o.TagIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentupload) MarshalJSON() ([]byte, error) {
    type Alias Documentupload

    if DocumentuploadMarshalled {
        return []byte("{}"), nil
    }
    DocumentuploadMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Workspace Domainentityref `json:"workspace"`
        
        Tags []string `json:"tags"`
        
        TagIds []string `json:"tagIds"`
        
        *Alias
    }{
        

        

        

        

        

        
        Tags: []string{""},
        

        

        
        TagIds: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}


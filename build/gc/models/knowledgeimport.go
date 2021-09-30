package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeimportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeimportDud struct { 
    Id string `json:"id"`


    


    


    


    


    Status string `json:"status"`


    Report Importreport `json:"report"`


    KnowledgeBase Knowledgebase `json:"knowledgeBase"`


    LanguageCode string `json:"languageCode"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Knowledgeimport
type Knowledgeimport struct { 
    


    // Name - Name of the import operation
    Name string `json:"name"`


    // UploadKey - Upload key
    UploadKey string `json:"uploadKey"`


    // FileType - file type of the document
    FileType string `json:"fileType"`


    // IgnoreHeaders - Ignore headers for the specified file
    IgnoreHeaders bool `json:"ignoreHeaders"`


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgeimport) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeimport) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeimport

    if KnowledgeimportMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeimportMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        UploadKey string `json:"uploadKey"`
        
        FileType string `json:"fileType"`
        
        IgnoreHeaders bool `json:"ignoreHeaders"`
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


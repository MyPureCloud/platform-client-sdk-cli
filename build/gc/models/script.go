package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScriptMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScriptDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Script
type Script struct { 
    


    // Name
    Name string `json:"name"`


    // VersionId
    VersionId string `json:"versionId"`


    // CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // PublishedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PublishedDate time.Time `json:"publishedDate"`


    // VersionDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    VersionDate time.Time `json:"versionDate"`


    // StartPageId
    StartPageId string `json:"startPageId"`


    // StartPageName
    StartPageName string `json:"startPageName"`


    // Features
    Features interface{} `json:"features"`


    // Variables
    Variables interface{} `json:"variables"`


    // CustomActions
    CustomActions interface{} `json:"customActions"`


    // Pages
    Pages []Page `json:"pages"`


    

}

// String returns a JSON representation of the model
func (o *Script) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Features = Interface{} 
    
    
    
     o.Variables = Interface{} 
    
    
    
     o.CustomActions = Interface{} 
    
    
    
     o.Pages = []Page{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Script) MarshalJSON() ([]byte, error) {
    type Alias Script

    if ScriptMarshalled {
        return []byte("{}"), nil
    }
    ScriptMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        VersionId string `json:"versionId"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        PublishedDate time.Time `json:"publishedDate"`
        
        VersionDate time.Time `json:"versionDate"`
        
        StartPageId string `json:"startPageId"`
        
        StartPageName string `json:"startPageName"`
        
        Features interface{} `json:"features"`
        
        Variables interface{} `json:"variables"`
        
        CustomActions interface{} `json:"customActions"`
        
        Pages []Page `json:"pages"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Features: Interface{},
        

        

        
        Variables: Interface{},
        

        

        
        CustomActions: Interface{},
        

        

        
        Pages: []Page{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}


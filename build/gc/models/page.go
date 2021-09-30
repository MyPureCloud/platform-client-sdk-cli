package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PageDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Page
type Page struct { 
    


    // Name
    Name string `json:"name"`


    // VersionId
    VersionId string `json:"versionId"`


    // CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // RootContainer
    RootContainer map[string]interface{} `json:"rootContainer"`


    // Properties
    Properties map[string]interface{} `json:"properties"`


    

}

// String returns a JSON representation of the model
func (o *Page) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.RootContainer = map[string]interface{}{"": Interface{}} 
    
    
    
     o.Properties = map[string]interface{}{"": Interface{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Page) MarshalJSON() ([]byte, error) {
    type Alias Page

    if PageMarshalled {
        return []byte("{}"), nil
    }
    PageMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        VersionId string `json:"versionId"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        RootContainer map[string]interface{} `json:"rootContainer"`
        
        Properties map[string]interface{} `json:"properties"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        RootContainer: map[string]interface{}{"": Interface{}},
        

        

        
        Properties: map[string]interface{}{"": Interface{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}


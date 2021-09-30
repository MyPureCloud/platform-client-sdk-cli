package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GkndocumentationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GkndocumentationresultDud struct { 
    


    


    


    

}

// Gkndocumentationresult
type Gkndocumentationresult struct { 
    // Content - The text or html content for the documentation entity. Will be returned in responses for certain entities.
    Content string `json:"content"`


    // Link - URL link for the documentation entity. Will be returned in responses for certain entities.
    Link string `json:"link"`


    // Title - The title of the documentation entity. Will be returned in responses for certain entities.
    Title string `json:"title"`


    // VarType - The search type. Will be returned in responses for certain entities.
    VarType string `json:"_type"`

}

// String returns a JSON representation of the model
func (o *Gkndocumentationresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Gkndocumentationresult) MarshalJSON() ([]byte, error) {
    type Alias Gkndocumentationresult

    if GkndocumentationresultMarshalled {
        return []byte("{}"), nil
    }
    GkndocumentationresultMarshalled = true

    return json.Marshal(&struct { 
        Content string `json:"content"`
        
        Link string `json:"link"`
        
        Title string `json:"title"`
        
        VarType string `json:"_type"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


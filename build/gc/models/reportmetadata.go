package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportmetadataDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Reportmetadata
type Reportmetadata struct { 
    


    // Name
    Name string `json:"name"`


    // Title
    Title string `json:"title"`


    // Description
    Description string `json:"description"`


    // Keywords
    Keywords []string `json:"keywords"`


    // AvailableLocales
    AvailableLocales []string `json:"availableLocales"`


    // Parameters
    Parameters []Parameter `json:"parameters"`


    // ExampleUrl
    ExampleUrl string `json:"exampleUrl"`


    

}

// String returns a JSON representation of the model
func (o *Reportmetadata) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Keywords = []string{""} 
    
    
    
     o.AvailableLocales = []string{""} 
    
    
    
     o.Parameters = []Parameter{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportmetadata) MarshalJSON() ([]byte, error) {
    type Alias Reportmetadata

    if ReportmetadataMarshalled {
        return []byte("{}"), nil
    }
    ReportmetadataMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        Keywords []string `json:"keywords"`
        
        AvailableLocales []string `json:"availableLocales"`
        
        Parameters []Parameter `json:"parameters"`
        
        ExampleUrl string `json:"exampleUrl"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Keywords: []string{""},
        

        

        
        AvailableLocales: []string{""},
        

        

        
        Parameters: []Parameter{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenerictemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenerictemplateDud struct { 
    


    


    


    


    

}

// Generictemplate
type Generictemplate struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Description - Text to show in the description.
    Description string `json:"description"`


    // Url - URL of an image.
    Url string `json:"url"`


    // Components - List of button components offered with this message content.
    Components []Recordingbuttoncomponent `json:"components"`


    // Actions - Actions to be taken.
    Actions Recordingcontentactions `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Generictemplate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Components = []Recordingbuttoncomponent{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generictemplate) MarshalJSON() ([]byte, error) {
    type Alias Generictemplate

    if GenerictemplateMarshalled {
        return []byte("{}"), nil
    }
    GenerictemplateMarshalled = true

    return json.Marshal(&struct { 
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        Url string `json:"url"`
        
        Components []Recordingbuttoncomponent `json:"components"`
        
        Actions Recordingcontentactions `json:"actions"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Components: []Recordingbuttoncomponent{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}


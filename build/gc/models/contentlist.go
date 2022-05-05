package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentlistDud struct { 
    


    


    


    


    


    


    

}

// Contentlist - List content object.
type Contentlist struct { 
    // Id - A unique ID assigned to this rich message content.
    Id string `json:"id"`


    // ListType - The type of list this instance represents.
    ListType string `json:"listType"`


    // Title - Text to show in the title.
    Title string `json:"title"`


    // Description - Text to show in the description.
    Description string `json:"description"`


    // SubmitLabel - Label for Submit button.
    SubmitLabel string `json:"submitLabel"`


    // Actions - The list actions (Deprecated).
    Actions Contentactions `json:"actions"`


    // Components - An array of component objects.
    Components []Listitemcomponent `json:"components"`

}

// String returns a JSON representation of the model
func (o *Contentlist) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Components = []Listitemcomponent{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentlist) MarshalJSON() ([]byte, error) {
    type Alias Contentlist

    if ContentlistMarshalled {
        return []byte("{}"), nil
    }
    ContentlistMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        ListType string `json:"listType"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        SubmitLabel string `json:"submitLabel"`
        
        Actions Contentactions `json:"actions"`
        
        Components []Listitemcomponent `json:"components"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Components: []Listitemcomponent{{}},
        

        
        Alias: (*Alias)(u),
    })
}


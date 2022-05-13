package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentgenericMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentgenericDud struct { 
    


    


    


    


    


    

}

// Contentgeneric - Deprecated, should use Card.
type Contentgeneric struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Description - Text to show in the description.
    Description string `json:"description"`


    // Image - URL of an image.
    Image string `json:"image"`


    // Video - URL of a video.
    Video string `json:"video"`


    // Actions - Actions to be taken (Deprecated).
    Actions Contentactions `json:"actions"`


    // Components - An array of component objects.
    Components []Buttoncomponent `json:"components"`

}

// String returns a JSON representation of the model
func (o *Contentgeneric) String() string {
    
    
    
    
    
     o.Components = []Buttoncomponent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentgeneric) MarshalJSON() ([]byte, error) {
    type Alias Contentgeneric

    if ContentgenericMarshalled {
        return []byte("{}"), nil
    }
    ContentgenericMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        Image string `json:"image"`
        
        Video string `json:"video"`
        
        Actions Contentactions `json:"actions"`
        
        Components []Buttoncomponent `json:"components"`
        *Alias
    }{

        


        


        


        


        


        
        Components: []Buttoncomponent{{}},
        

        Alias: (*Alias)(u),
    })
}


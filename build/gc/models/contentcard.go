package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentcardMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentcardDud struct { 
    


    


    


    


    


    

}

// Contentcard - Card content object.
type Contentcard struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Description - Text to show in the description.
    Description string `json:"description"`


    // Image - URL of an image.
    Image string `json:"image"`


    // Video - URL of a video.
    Video string `json:"video"`


    // DefaultAction - The default button action.
    DefaultAction Contentcardaction `json:"defaultAction"`


    // Actions - An array of action objects.
    Actions []Contentcardaction `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Contentcard) String() string {
    
    
    
    
    
     o.Actions = []Contentcardaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentcard) MarshalJSON() ([]byte, error) {
    type Alias Contentcard

    if ContentcardMarshalled {
        return []byte("{}"), nil
    }
    ContentcardMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        Image string `json:"image"`
        
        Video string `json:"video"`
        
        DefaultAction Contentcardaction `json:"defaultAction"`
        
        Actions []Contentcardaction `json:"actions"`
        *Alias
    }{

        


        


        


        


        


        
        Actions: []Contentcardaction{{}},
        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormlistpickeritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormlistpickeritemDud struct { 
    


    


    

}

// Formlistpickeritem - An item in a list picker section
type Formlistpickeritem struct { 
    // Id - Unique identifier for the item
    Id string `json:"id"`


    // Title - Title of the item
    Title string `json:"title"`


    // ImageUrl - URL of the image to display
    ImageUrl string `json:"imageUrl"`

}

// String returns a JSON representation of the model
func (o *Formlistpickeritem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formlistpickeritem) MarshalJSON() ([]byte, error) {
    type Alias Formlistpickeritem

    if FormlistpickeritemMarshalled {
        return []byte("{}"), nil
    }
    FormlistpickeritemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        ImageUrl string `json:"imageUrl"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListpickeritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListpickeritemDud struct { 
    


    


    


    

}

// Listpickeritem
type Listpickeritem struct { 
    // Id - Unique identifier for the list picker item
    Id string `json:"id"`


    // Subtitle - Additional text providing more details about the item.
    Subtitle string `json:"subtitle"`


    // Title - The main text displayed for the item.
    Title string `json:"title"`


    // Selected - Whether the item is selected.
    Selected bool `json:"selected"`

}

// String returns a JSON representation of the model
func (o *Listpickeritem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listpickeritem) MarshalJSON() ([]byte, error) {
    type Alias Listpickeritem

    if ListpickeritemMarshalled {
        return []byte("{}"), nil
    }
    ListpickeritemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Subtitle string `json:"subtitle"`
        
        Title string `json:"title"`
        
        Selected bool `json:"selected"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


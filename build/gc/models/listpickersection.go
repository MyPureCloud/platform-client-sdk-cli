package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListpickersectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListpickersectionDud struct { 
    


    


    

}

// Listpickersection
type Listpickersection struct { 
    // Items - List of items to choose from in the section
    Items []Listpickeritem `json:"items"`


    // MultipleSelection - Whether multiple items can be selected in this section.
    MultipleSelection bool `json:"multipleSelection"`


    // Title - Required title for the section.
    Title string `json:"title"`

}

// String returns a JSON representation of the model
func (o *Listpickersection) String() string {
     o.Items = []Listpickeritem{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listpickersection) MarshalJSON() ([]byte, error) {
    type Alias Listpickersection

    if ListpickersectionMarshalled {
        return []byte("{}"), nil
    }
    ListpickersectionMarshalled = true

    return json.Marshal(&struct {
        
        Items []Listpickeritem `json:"items"`
        
        MultipleSelection bool `json:"multipleSelection"`
        
        Title string `json:"title"`
        *Alias
    }{

        
        Items: []Listpickeritem{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}


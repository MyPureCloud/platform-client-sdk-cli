package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormlistpickersectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormlistpickersectionDud struct { 
    


    


    

}

// Formlistpickersection - Section within a list picker
type Formlistpickersection struct { 
    // Title - Title of the section
    Title string `json:"title"`


    // MultipleSelection - Whether multiple items can be selected
    MultipleSelection bool `json:"multipleSelection"`


    // Items - Items in this section
    Items []Formlistpickeritem `json:"items"`

}

// String returns a JSON representation of the model
func (o *Formlistpickersection) String() string {
    
    
     o.Items = []Formlistpickeritem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formlistpickersection) MarshalJSON() ([]byte, error) {
    type Alias Formlistpickersection

    if FormlistpickersectionMarshalled {
        return []byte("{}"), nil
    }
    FormlistpickersectionMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        MultipleSelection bool `json:"multipleSelection"`
        
        Items []Formlistpickeritem `json:"items"`
        *Alias
    }{

        


        


        
        Items: []Formlistpickeritem{{}},
        

        Alias: (*Alias)(u),
    })
}


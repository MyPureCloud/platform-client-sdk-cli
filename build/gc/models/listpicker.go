package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListpickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListpickerDud struct { 
    


    


    

}

// Listpicker
type Listpicker struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the description.
    Subtitle string `json:"subtitle"`


    // Sections - An array of sections in the List Picker.
    Sections []Listpickersection `json:"sections"`

}

// String returns a JSON representation of the model
func (o *Listpicker) String() string {
    
    
     o.Sections = []Listpickersection{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listpicker) MarshalJSON() ([]byte, error) {
    type Alias Listpicker

    if ListpickerMarshalled {
        return []byte("{}"), nil
    }
    ListpickerMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        Sections []Listpickersection `json:"sections"`
        *Alias
    }{

        


        


        
        Sections: []Listpickersection{{}},
        

        Alias: (*Alias)(u),
    })
}


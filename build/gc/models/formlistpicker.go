package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormlistpickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormlistpickerDud struct { 
    


    

}

// Formlistpicker - List picker component configuration
type Formlistpicker struct { 
    // Id - Unique identifier for the list picker
    Id string `json:"id"`


    // Sections - Sections in the list picker
    Sections []Formlistpickersection `json:"sections"`

}

// String returns a JSON representation of the model
func (o *Formlistpicker) String() string {
    
     o.Sections = []Formlistpickersection{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formlistpicker) MarshalJSON() ([]byte, error) {
    type Alias Formlistpicker

    if FormlistpickerMarshalled {
        return []byte("{}"), nil
    }
    FormlistpickerMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Sections []Formlistpickersection `json:"sections"`
        *Alias
    }{

        


        
        Sections: []Formlistpickersection{{}},
        

        Alias: (*Alias)(u),
    })
}


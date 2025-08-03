package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormdatepickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormdatepickerDud struct { 
    


    


    


    

}

// Formdatepicker - Date picker component for form input
type Formdatepicker struct { 
    // Id - Unique identifier for the date picker
    Id string `json:"id"`


    // Title - Title of the date picker
    Title string `json:"title"`


    // Subtitle - Subtitle of the date picker
    Subtitle string `json:"subtitle"`


    // DateDisplayFormat - Date display format
    DateDisplayFormat string `json:"dateDisplayFormat"`

}

// String returns a JSON representation of the model
func (o *Formdatepicker) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formdatepicker) MarshalJSON() ([]byte, error) {
    type Alias Formdatepicker

    if FormdatepickerMarshalled {
        return []byte("{}"), nil
    }
    FormdatepickerMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        DateDisplayFormat string `json:"dateDisplayFormat"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


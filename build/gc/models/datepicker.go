package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DatepickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DatepickerDud struct { 
    


    


    

}

// Datepicker
type Datepicker struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the description.
    Subtitle string `json:"subtitle"`


    // DatePickerAvailableDateTimes - An array of available times objects.
    DatePickerAvailableDateTimes []Datepickeravailabledatetime `json:"datePickerAvailableDateTimes"`

}

// String returns a JSON representation of the model
func (o *Datepicker) String() string {
    
    
     o.DatePickerAvailableDateTimes = []Datepickeravailabledatetime{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Datepicker) MarshalJSON() ([]byte, error) {
    type Alias Datepicker

    if DatepickerMarshalled {
        return []byte("{}"), nil
    }
    DatepickerMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        DatePickerAvailableDateTimes []Datepickeravailabledatetime `json:"datePickerAvailableDateTimes"`
        *Alias
    }{

        


        


        
        DatePickerAvailableDateTimes: []Datepickeravailabledatetime{{}},
        

        Alias: (*Alias)(u),
    })
}


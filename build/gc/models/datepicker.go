package models
import (
    "time"
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
    // Id - Optional unique identifier to help map component replies to form messages where multiple DatePickers can be present.
    Id string `json:"id"`


    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the description.
    Subtitle string `json:"subtitle"`


    // DatePickerAvailableDateTimes - An array of available times objects.
    DatePickerAvailableDateTimes []Datepickeravailabledatetime `json:"datePickerAvailableDateTimes"`


    // DateSelected - Selected date response from end customer. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateSelected time.Time `json:"dateSelected"`

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
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        DatePickerAvailableDateTimes []Datepickeravailabledatetime `json:"datePickerAvailableDateTimes"`
        
        DateSelected time.Time `json:"dateSelected"`
        *Alias
    }{

        


        


        


        
        DatePickerAvailableDateTimes: []Datepickeravailabledatetime{{}},
        


        

        Alias: (*Alias)(u),
    })
}


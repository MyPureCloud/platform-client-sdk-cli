package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentdatepickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentdatepickerDud struct { 
    


    


    


    


    


    

}

// Contentdatepicker - DatePicker content object.
type Contentdatepicker struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the description.
    Subtitle string `json:"subtitle"`


    // ImageUrl - URL of an image
    ImageUrl string `json:"imageUrl"`


    // DateMinimum - The minimum Date Enabled in the datepicker calendar. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateMinimum time.Time `json:"dateMinimum"`


    // DateMaximum - The maximum Date Enabled in the datepicker calendar. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateMaximum time.Time `json:"dateMaximum"`


    // AvailableTimes - An array of available times objects.
    AvailableTimes []Contentdatepickeravailabletime `json:"availableTimes"`

}

// String returns a JSON representation of the model
func (o *Contentdatepicker) String() string {
    
    
    
    
    
     o.AvailableTimes = []Contentdatepickeravailabletime{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentdatepicker) MarshalJSON() ([]byte, error) {
    type Alias Contentdatepicker

    if ContentdatepickerMarshalled {
        return []byte("{}"), nil
    }
    ContentdatepickerMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        ImageUrl string `json:"imageUrl"`
        
        DateMinimum time.Time `json:"dateMinimum"`
        
        DateMaximum time.Time `json:"dateMaximum"`
        
        AvailableTimes []Contentdatepickeravailabletime `json:"availableTimes"`
        *Alias
    }{

        


        


        


        


        


        
        AvailableTimes: []Contentdatepickeravailabletime{{}},
        

        Alias: (*Alias)(u),
    })
}


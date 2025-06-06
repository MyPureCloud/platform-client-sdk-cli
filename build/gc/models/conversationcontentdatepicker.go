package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentdatepickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentdatepickerDud struct { 
    


    


    


    


    


    


    


    


    

}

// Conversationcontentdatepicker - DateTimePicker content object.
type Conversationcontentdatepicker struct { 
    // Id - Optional unique identifier to help map component replies to form messages where multiple DatePickers can be present.
    Id string `json:"id"`


    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the description.
    Subtitle string `json:"subtitle"`


    // ImageUrl - URL of an image
    ImageUrl string `json:"imageUrl"`


    // DateMinimum - The minimum Date Enabled in the datepicker calendar, format: ISO 8601.
    DateMinimum time.Time `json:"dateMinimum"`


    // DateMaximum - The maximum Date Enabled in the datepicker calendar, format: ISO 8601.
    DateMaximum time.Time `json:"dateMaximum"`


    // Location - Location of the event.
    Location Conversationcontentlocation `json:"location"`


    // AvailableTimes - An array of available times objects.
    AvailableTimes []Conversationcontentdatepickeravailabletime `json:"availableTimes"`


    // DateDisplayFormat - The format the date should be presented to the end user.
    DateDisplayFormat string `json:"dateDisplayFormat"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentdatepicker) String() string {
    
    
    
    
    
    
    
     o.AvailableTimes = []Conversationcontentdatepickeravailabletime{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentdatepicker) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentdatepicker

    if ConversationcontentdatepickerMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentdatepickerMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        ImageUrl string `json:"imageUrl"`
        
        DateMinimum time.Time `json:"dateMinimum"`
        
        DateMaximum time.Time `json:"dateMaximum"`
        
        Location Conversationcontentlocation `json:"location"`
        
        AvailableTimes []Conversationcontentdatepickeravailabletime `json:"availableTimes"`
        
        DateDisplayFormat string `json:"dateDisplayFormat"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        AvailableTimes: []Conversationcontentdatepickeravailabletime{{}},
        


        

        Alias: (*Alias)(u),
    })
}


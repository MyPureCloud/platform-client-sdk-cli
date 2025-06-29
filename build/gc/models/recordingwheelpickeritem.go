package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingwheelpickeritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingwheelpickeritemDud struct { 
    


    


    


    

}

// Recordingwheelpickeritem
type Recordingwheelpickeritem struct { 
    // Id - Unique identifier for the wheel picker item.
    Id string `json:"id"`


    // Title - The main text displayed for the item.
    Title string `json:"title"`


    // Value - The value of the item
    Value string `json:"value"`


    // IsSelected - Indicates if the item is selected by end customer.
    IsSelected bool `json:"isSelected"`

}

// String returns a JSON representation of the model
func (o *Recordingwheelpickeritem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingwheelpickeritem) MarshalJSON() ([]byte, error) {
    type Alias Recordingwheelpickeritem

    if RecordingwheelpickeritemMarshalled {
        return []byte("{}"), nil
    }
    RecordingwheelpickeritemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Value string `json:"value"`
        
        IsSelected bool `json:"isSelected"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


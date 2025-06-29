package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingwheelpickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingwheelpickerDud struct { 
    


    

}

// Recordingwheelpicker
type Recordingwheelpicker struct { 
    // Id - Optional unique identifier to help map component replies to form messages where multiple Wheel Pickers can be present.
    Id string `json:"id"`


    // Items - An array of options in the Wheel Picker.
    Items []Recordingwheelpickeritem `json:"items"`

}

// String returns a JSON representation of the model
func (o *Recordingwheelpicker) String() string {
    
     o.Items = []Recordingwheelpickeritem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingwheelpicker) MarshalJSON() ([]byte, error) {
    type Alias Recordingwheelpicker

    if RecordingwheelpickerMarshalled {
        return []byte("{}"), nil
    }
    RecordingwheelpickerMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Items []Recordingwheelpickeritem `json:"items"`
        *Alias
    }{

        


        
        Items: []Recordingwheelpickeritem{{}},
        

        Alias: (*Alias)(u),
    })
}


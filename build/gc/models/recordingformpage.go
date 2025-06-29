package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingformpageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingformpageDud struct { 
    


    


    

}

// Recordingformpage
type Recordingformpage struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the subtitle.
    Subtitle string `json:"subtitle"`


    // PageComponents - Page components in this form page.
    PageComponents []Recordingformpagecomponent `json:"pageComponents"`

}

// String returns a JSON representation of the model
func (o *Recordingformpage) String() string {
    
    
     o.PageComponents = []Recordingformpagecomponent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingformpage) MarshalJSON() ([]byte, error) {
    type Alias Recordingformpage

    if RecordingformpageMarshalled {
        return []byte("{}"), nil
    }
    RecordingformpageMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        PageComponents []Recordingformpagecomponent `json:"pageComponents"`
        *Alias
    }{

        


        


        
        PageComponents: []Recordingformpagecomponent{{}},
        

        Alias: (*Alias)(u),
    })
}


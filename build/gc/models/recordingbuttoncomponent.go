package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingbuttoncomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingbuttoncomponentDud struct { 
    


    


    

}

// Recordingbuttoncomponent - Structured template button object.
type Recordingbuttoncomponent struct { 
    // Title
    Title string `json:"title"`


    // Actions
    Actions Recordingcontentactions `json:"actions"`


    // IsSelected
    IsSelected bool `json:"isSelected"`

}

// String returns a JSON representation of the model
func (o *Recordingbuttoncomponent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingbuttoncomponent) MarshalJSON() ([]byte, error) {
    type Alias Recordingbuttoncomponent

    if RecordingbuttoncomponentMarshalled {
        return []byte("{}"), nil
    }
    RecordingbuttoncomponentMarshalled = true

    return json.Marshal(&struct { 
        Title string `json:"title"`
        
        Actions Recordingcontentactions `json:"actions"`
        
        IsSelected bool `json:"isSelected"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


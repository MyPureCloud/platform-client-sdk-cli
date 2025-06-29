package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingformpagecomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingformpagecomponentDud struct { 
    


    


    


    


    

}

// Recordingformpagecomponent
type Recordingformpagecomponent struct { 
    // FormComponentType - Type of this form component element.
    FormComponentType string `json:"formComponentType"`


    // DatePicker - Date Picker content.
    DatePicker Datepicker `json:"datePicker"`


    // WheelPicker - Wheel Picker content.
    WheelPicker Recordingwheelpicker `json:"wheelPicker"`


    // ListPicker - List Picker content.
    ListPicker Listpicker `json:"listPicker"`


    // Input - Input content.
    Input Recordinginput `json:"input"`

}

// String returns a JSON representation of the model
func (o *Recordingformpagecomponent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingformpagecomponent) MarshalJSON() ([]byte, error) {
    type Alias Recordingformpagecomponent

    if RecordingformpagecomponentMarshalled {
        return []byte("{}"), nil
    }
    RecordingformpagecomponentMarshalled = true

    return json.Marshal(&struct {
        
        FormComponentType string `json:"formComponentType"`
        
        DatePicker Datepicker `json:"datePicker"`
        
        WheelPicker Recordingwheelpicker `json:"wheelPicker"`
        
        ListPicker Listpicker `json:"listPicker"`
        
        Input Recordinginput `json:"input"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


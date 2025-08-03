package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormpagecomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormpagecomponentDud struct { 
    


    


    


    


    

}

// Formpagecomponent - A component within a form page
type Formpagecomponent struct { 
    // FormComponentType - Type of the component
    FormComponentType string `json:"formComponentType"`


    // ListPicker - List picker configuration
    ListPicker Formlistpicker `json:"listPicker"`


    // DatePicker - Date picker configuration
    DatePicker Formdatepicker `json:"datePicker"`


    // Input - Input field configuration
    Input Input `json:"input"`


    // WheelPicker - Wheel picker configuration
    WheelPicker Wheelpicker `json:"wheelPicker"`

}

// String returns a JSON representation of the model
func (o *Formpagecomponent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formpagecomponent) MarshalJSON() ([]byte, error) {
    type Alias Formpagecomponent

    if FormpagecomponentMarshalled {
        return []byte("{}"), nil
    }
    FormpagecomponentMarshalled = true

    return json.Marshal(&struct {
        
        FormComponentType string `json:"formComponentType"`
        
        ListPicker Formlistpicker `json:"listPicker"`
        
        DatePicker Formdatepicker `json:"datePicker"`
        
        Input Input `json:"input"`
        
        WheelPicker Wheelpicker `json:"wheelPicker"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


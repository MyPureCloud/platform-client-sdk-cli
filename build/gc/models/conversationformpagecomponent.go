package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationformpagecomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationformpagecomponentDud struct { 
    


    


    


    


    

}

// Conversationformpagecomponent - Form component element. Examples include: List Picker, Date Picker, Wheel Picker and Input text.
type Conversationformpagecomponent struct { 
    // FormComponentType - Type of this form component element
    FormComponentType string `json:"formComponentType"`


    // DatePicker - Date Picker content.
    DatePicker Conversationcontentdatepicker `json:"datePicker"`


    // WheelPicker - Wheel Picker content.
    WheelPicker Conversationcontentwheelpicker `json:"wheelPicker"`


    // ListPicker - List Picker content.
    ListPicker Conversationcontentlistpicker `json:"listPicker"`


    // Input - Input content.
    Input Conversationcontentinput `json:"input"`

}

// String returns a JSON representation of the model
func (o *Conversationformpagecomponent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationformpagecomponent) MarshalJSON() ([]byte, error) {
    type Alias Conversationformpagecomponent

    if ConversationformpagecomponentMarshalled {
        return []byte("{}"), nil
    }
    ConversationformpagecomponentMarshalled = true

    return json.Marshal(&struct {
        
        FormComponentType string `json:"formComponentType"`
        
        DatePicker Conversationcontentdatepicker `json:"datePicker"`
        
        WheelPicker Conversationcontentwheelpicker `json:"wheelPicker"`
        
        ListPicker Conversationcontentlistpicker `json:"listPicker"`
        
        Input Conversationcontentinput `json:"input"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


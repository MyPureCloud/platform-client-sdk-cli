package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormstracktriggerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormstracktriggerDud struct { 
    


    


    


    

}

// Formstracktrigger - Details about a forms tracking event trigger
type Formstracktrigger struct { 
    // Selector - Form element that triggers the form submitted or abandoned event.
    Selector string `json:"selector"`


    // FormName - Prefix for the form submitted or abandoned event name.
    FormName string `json:"formName"`


    // CaptureDataOnFormAbandon - Whether to capture the form data in the form abandoned event.
    CaptureDataOnFormAbandon bool `json:"captureDataOnFormAbandon"`


    // CaptureDataOnFormSubmit - Whether to capture the form data in the form submitted event.
    CaptureDataOnFormSubmit bool `json:"captureDataOnFormSubmit"`

}

// String returns a JSON representation of the model
func (o *Formstracktrigger) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formstracktrigger) MarshalJSON() ([]byte, error) {
    type Alias Formstracktrigger

    if FormstracktriggerMarshalled {
        return []byte("{}"), nil
    }
    FormstracktriggerMarshalled = true

    return json.Marshal(&struct {
        
        Selector string `json:"selector"`
        
        FormName string `json:"formName"`
        
        CaptureDataOnFormAbandon bool `json:"captureDataOnFormAbandon"`
        
        CaptureDataOnFormSubmit bool `json:"captureDataOnFormSubmit"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


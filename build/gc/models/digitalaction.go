package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DigitalactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DigitalactionDud struct { 
    


    


    


    


    


    

}

// Digitalaction
type Digitalaction struct { 
    // UpdateContactColumnActionSettings - The settings for an 'update contact column' action.
    UpdateContactColumnActionSettings Updatecontactcolumnactionsettings `json:"updateContactColumnActionSettings"`


    // DoNotSendActionSettings - The settings for a 'do not send' action.
    DoNotSendActionSettings interface{} `json:"doNotSendActionSettings"`


    // AppendToDncActionSettings - The settings for an 'Append to DNC' action.
    AppendToDncActionSettings Appendtodncactionsettings `json:"appendToDncActionSettings"`


    // MarkContactUncontactableActionSettings - The settings for a 'mark contact uncontactable' action.
    MarkContactUncontactableActionSettings Markcontactuncontactableactionsettings `json:"markContactUncontactableActionSettings"`


    // MarkContactAddressUncontactableActionSettings - The settings for an 'mark contact address uncontactable' action.
    MarkContactAddressUncontactableActionSettings interface{} `json:"markContactAddressUncontactableActionSettings"`


    // SetContentTemplateActionSettings - The settings for a 'Set content template' action.
    SetContentTemplateActionSettings Setcontenttemplateactionsettings `json:"setContentTemplateActionSettings"`

}

// String returns a JSON representation of the model
func (o *Digitalaction) String() string {
    
     o.DoNotSendActionSettings = Interface{} 
    
    
     o.MarkContactAddressUncontactableActionSettings = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Digitalaction) MarshalJSON() ([]byte, error) {
    type Alias Digitalaction

    if DigitalactionMarshalled {
        return []byte("{}"), nil
    }
    DigitalactionMarshalled = true

    return json.Marshal(&struct {
        
        UpdateContactColumnActionSettings Updatecontactcolumnactionsettings `json:"updateContactColumnActionSettings"`
        
        DoNotSendActionSettings interface{} `json:"doNotSendActionSettings"`
        
        AppendToDncActionSettings Appendtodncactionsettings `json:"appendToDncActionSettings"`
        
        MarkContactUncontactableActionSettings Markcontactuncontactableactionsettings `json:"markContactUncontactableActionSettings"`
        
        MarkContactAddressUncontactableActionSettings interface{} `json:"markContactAddressUncontactableActionSettings"`
        
        SetContentTemplateActionSettings Setcontenttemplateactionsettings `json:"setContentTemplateActionSettings"`
        *Alias
    }{

        


        
        DoNotSendActionSettings: Interface{},
        


        


        


        
        MarkContactAddressUncontactableActionSettings: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}


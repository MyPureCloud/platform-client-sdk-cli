package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialeractionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialeractionDud struct { 
    


    


    


    


    


    


    


    


    

}

// Dialeraction
type Dialeraction struct { 
    // VarType - The type of this DialerAction.
    VarType string `json:"type"`


    // ActionTypeName - Additional type specification for this DialerAction.
    ActionTypeName string `json:"actionTypeName"`


    // UpdateOption - Specifies how a contact attribute should be updated. Required for MODIFY_CONTACT_ATTRIBUTE.
    UpdateOption string `json:"updateOption"`


    // Properties - A map of key-value pairs pertinent to the DialerAction. Different types of DialerActions require different properties. MODIFY_CONTACT_ATTRIBUTE with an updateOption of SET takes a contact column as the key and accepts any value. SCHEDULE_CALLBACK takes a key 'callbackOffset' that specifies how far in the future the callback should be scheduled, in minutes. SET_CALLER_ID takes two keys: 'callerAddress', which should be the caller id phone number, and 'callerName'. For either key, you can also specify a column on the contact to get the value from. To do this, specify 'contact.Column', where 'Column' is the name of the contact column from which to get the value. SET_SKILLS takes a key 'skills' with an array of skill ids wrapped into a string (Example: {'skills': '['skillIdHere']'} ).
    Properties map[string]string `json:"properties"`


    // DataAction - The Data Action to use for this action. Required for a dataActionBehavior.
    DataAction Domainentityref `json:"dataAction"`


    // ContactColumnToDataActionFieldMappings - A list of mappings defining which contact data fields will be passed to which data action input fields for this condition. Valid for a dataActionBehavior.
    ContactColumnToDataActionFieldMappings []Contactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings"`


    // ContactIdField - The input field from the data action that the contactId will be passed to for this condition. Valid for a dataActionBehavior.
    ContactIdField string `json:"contactIdField"`


    // CallAnalysisResultField - The input field from the data action that the callAnalysisResult will be passed to for this condition. Valid for a wrapup dataActionBehavior.
    CallAnalysisResultField string `json:"callAnalysisResultField"`


    // AgentWrapupField - The input field from the data action that the agentWrapup will be passed to for this condition. Valid for a wrapup dataActionBehavior.
    AgentWrapupField string `json:"agentWrapupField"`

}

// String returns a JSON representation of the model
func (o *Dialeraction) String() string {
    
    
    
     o.Properties = map[string]string{"": ""} 
    
     o.ContactColumnToDataActionFieldMappings = []Contactcolumntodataactionfieldmapping{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialeraction) MarshalJSON() ([]byte, error) {
    type Alias Dialeraction

    if DialeractionMarshalled {
        return []byte("{}"), nil
    }
    DialeractionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        ActionTypeName string `json:"actionTypeName"`
        
        UpdateOption string `json:"updateOption"`
        
        Properties map[string]string `json:"properties"`
        
        DataAction Domainentityref `json:"dataAction"`
        
        ContactColumnToDataActionFieldMappings []Contactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings"`
        
        ContactIdField string `json:"contactIdField"`
        
        CallAnalysisResultField string `json:"callAnalysisResultField"`
        
        AgentWrapupField string `json:"agentWrapupField"`
        *Alias
    }{

        


        


        


        
        Properties: map[string]string{"": ""},
        


        


        
        ContactColumnToDataActionFieldMappings: []Contactcolumntodataactionfieldmapping{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}


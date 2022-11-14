package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataactionconditionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataactionconditionsettingsDud struct { 
    


    


    


    


    

}

// Dataactionconditionsettings
type Dataactionconditionsettings struct { 
    // DataActionId - The Data Action Id to use for this condition.
    DataActionId string `json:"dataActionId"`


    // ContactIdField - The input field from the data action that the contactId will be passed into.
    ContactIdField string `json:"contactIdField"`


    // DataNotFoundResolution - The result of this condition if the data action returns a result indicating there was no data.
    DataNotFoundResolution bool `json:"dataNotFoundResolution"`


    // Predicates - A list of predicates defining the comparisons to use for this condition.
    Predicates []Digitaldataactionconditionpredicate `json:"predicates"`


    // ContactColumnToDataActionFieldMappings - A list of mappings defining which contact data fields will be passed to which data action input fields.
    ContactColumnToDataActionFieldMappings []Dataactioncontactcolumnfieldmapping `json:"contactColumnToDataActionFieldMappings"`

}

// String returns a JSON representation of the model
func (o *Dataactionconditionsettings) String() string {
    
    
    
     o.Predicates = []Digitaldataactionconditionpredicate{{}} 
     o.ContactColumnToDataActionFieldMappings = []Dataactioncontactcolumnfieldmapping{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataactionconditionsettings) MarshalJSON() ([]byte, error) {
    type Alias Dataactionconditionsettings

    if DataactionconditionsettingsMarshalled {
        return []byte("{}"), nil
    }
    DataactionconditionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        DataActionId string `json:"dataActionId"`
        
        ContactIdField string `json:"contactIdField"`
        
        DataNotFoundResolution bool `json:"dataNotFoundResolution"`
        
        Predicates []Digitaldataactionconditionpredicate `json:"predicates"`
        
        ContactColumnToDataActionFieldMappings []Dataactioncontactcolumnfieldmapping `json:"contactColumnToDataActionFieldMappings"`
        *Alias
    }{

        


        


        


        
        Predicates: []Digitaldataactionconditionpredicate{{}},
        


        
        ContactColumnToDataActionFieldMappings: []Dataactioncontactcolumnfieldmapping{{}},
        

        Alias: (*Alias)(u),
    })
}


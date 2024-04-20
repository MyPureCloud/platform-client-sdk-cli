package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportedlanguagesinfodefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportedlanguagesinfodefinitionDud struct { 
    


    


    


    


    

}

// Supportedlanguagesinfodefinition
type Supportedlanguagesinfodefinition struct { 
    // Language - The given supported Language
    Language string `json:"language"`


    // IntentClassification - The boolean status of if intent classification is supported in this language
    IntentClassification bool `json:"intentClassification"`


    // Status - The language release status
    Status string `json:"status"`


    // SupportedEntityTypes - The supported entity types for this language
    SupportedEntityTypes []string `json:"supportedEntityTypes"`


    // SupportedEntityTypeConfiguration - The configuration for the supported entity types
    SupportedEntityTypeConfiguration Supportedentitytypestatus `json:"supportedEntityTypeConfiguration"`

}

// String returns a JSON representation of the model
func (o *Supportedlanguagesinfodefinition) String() string {
    
    
    
     o.SupportedEntityTypes = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportedlanguagesinfodefinition) MarshalJSON() ([]byte, error) {
    type Alias Supportedlanguagesinfodefinition

    if SupportedlanguagesinfodefinitionMarshalled {
        return []byte("{}"), nil
    }
    SupportedlanguagesinfodefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Language string `json:"language"`
        
        IntentClassification bool `json:"intentClassification"`
        
        Status string `json:"status"`
        
        SupportedEntityTypes []string `json:"supportedEntityTypes"`
        
        SupportedEntityTypeConfiguration Supportedentitytypestatus `json:"supportedEntityTypeConfiguration"`
        *Alias
    }{

        


        


        


        
        SupportedEntityTypes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}


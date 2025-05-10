package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactbulksearchparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactbulksearchparametersDud struct { 
    


    


    

}

// Contactbulksearchparameters
type Contactbulksearchparameters struct { 
    // ContactListFilterId - Contact List Filter ID. Either this property or criteria is required.
    ContactListFilterId string `json:"contactListFilterId"`


    // Criteria - Criteria to filter the contacts by. Either this property or contactListFilterId is required.
    Criteria Contactbulksearchcriteria `json:"criteria"`


    // GenerateDownloadURI - Whether to do backup export as part of Bulk Operation or not. Default: true.
    GenerateDownloadURI bool `json:"generateDownloadURI"`

}

// String returns a JSON representation of the model
func (o *Contactbulksearchparameters) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactbulksearchparameters) MarshalJSON() ([]byte, error) {
    type Alias Contactbulksearchparameters

    if ContactbulksearchparametersMarshalled {
        return []byte("{}"), nil
    }
    ContactbulksearchparametersMarshalled = true

    return json.Marshal(&struct {
        
        ContactListFilterId string `json:"contactListFilterId"`
        
        Criteria Contactbulksearchcriteria `json:"criteria"`
        
        GenerateDownloadURI bool `json:"generateDownloadURI"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}


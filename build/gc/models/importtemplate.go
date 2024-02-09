package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImporttemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImporttemplateDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    


    


    ImportStatus Importstatus `json:"importStatus"`


    SelfUri string `json:"selfUri"`

}

// Importtemplate
type Importtemplate struct { 
    


    // Name - The name of the import template.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // ContactListTemplate - ContactListTemplate for this ImportTemplate.
    ContactListTemplate Domainentityref `json:"contactListTemplate"`


    // ContactListFilter - ContactListFilter for this ImportTemplate.
    ContactListFilter Domainentityref `json:"contactListFilter"`


    // UseSplittingCriteria - Whether or not to use splitting criteria. Default is false.
    UseSplittingCriteria bool `json:"useSplittingCriteria"`


    // SplittingInformation - How to split contact records, required if useSplittingCriteria is true.
    SplittingInformation Splittinginformation `json:"splittingInformation"`


    // ListNameFormat - The list name format for target ContactLists. When Custom is provided, customListNameFormatValue is required.
    ListNameFormat string `json:"listNameFormat"`


    // CustomListNameFormatValue - Custom value for the list name format, at least %N is required. Any character other than the specified tokens will be used as is. Available tokens: %N: ListNamePrefix; %P: Part number; %F: Filter name; %C: Column value; YYYY: year; MM: month; DD: day; hh: hour; mm: minute; ss: second.
    CustomListNameFormatValue string `json:"customListNameFormatValue"`


    


    

}

// String returns a JSON representation of the model
func (o *Importtemplate) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importtemplate) MarshalJSON() ([]byte, error) {
    type Alias Importtemplate

    if ImporttemplateMarshalled {
        return []byte("{}"), nil
    }
    ImporttemplateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        ContactListTemplate Domainentityref `json:"contactListTemplate"`
        
        ContactListFilter Domainentityref `json:"contactListFilter"`
        
        UseSplittingCriteria bool `json:"useSplittingCriteria"`
        
        SplittingInformation Splittinginformation `json:"splittingInformation"`
        
        ListNameFormat string `json:"listNameFormat"`
        
        CustomListNameFormatValue string `json:"customListNameFormatValue"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


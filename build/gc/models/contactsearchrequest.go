package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsearchrequestDud struct { 
    


    


    


    


    

}

// Contactsearchrequest
type Contactsearchrequest struct { 
    // PageNumber - Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)
    PageNumber int `json:"pageNumber"`


    // PageSize - Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)
    PageSize int `json:"pageSize"`


    // DivisionIds - Which divisions to search, up to 50
    DivisionIds []string `json:"divisionIds"`


    // Expand - Which fields, if any, to expand
    Expand []string `json:"expand"`


    // Operation - Search operation to execute, currently supports {@code simpleSearch} only.
    Operation Contactsearchoperation `json:"operation"`

}

// String returns a JSON representation of the model
func (o *Contactsearchrequest) String() string {
    
    
     o.DivisionIds = []string{""} 
     o.Expand = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactsearchrequest

    if ContactsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        DivisionIds []string `json:"divisionIds"`
        
        Expand []string `json:"expand"`
        
        Operation Contactsearchoperation `json:"operation"`
        *Alias
    }{

        


        


        
        DivisionIds: []string{""},
        


        
        Expand: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}


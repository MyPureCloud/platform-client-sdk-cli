package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsimplesearchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsimplesearchDud struct { 
    


    


    

}

// Contactsimplesearch
type Contactsimplesearch struct { 
    // Query - User supplied search keywords (no special syntax is currently supported)
    Query string `json:"query"`


    // SortOrder - The External Contact field to sort by. Any of: [firstName, lastName, middleName, title]. Direction: [asc, desc]. e.g. \"firstName:asc\", \"title:desc\"
    SortOrder []string `json:"sortOrder"`


    // Ids - List of External Contact ids to exact match in search result. Optional filter, up to 100
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Contactsimplesearch) String() string {
    
     o.SortOrder = []string{""} 
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsimplesearch) MarshalJSON() ([]byte, error) {
    type Alias Contactsimplesearch

    if ContactsimplesearchMarshalled {
        return []byte("{}"), nil
    }
    ContactsimplesearchMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        SortOrder []string `json:"sortOrder"`
        
        Ids []string `json:"ids"`
        *Alias
    }{

        


        
        SortOrder: []string{""},
        


        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}


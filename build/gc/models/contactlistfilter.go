package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistfilterDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contactlistfilter
type Contactlistfilter struct { 
    


    // Name - The name of the list.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // ContactList - The contact list the filter is based on.
    ContactList Domainentityref `json:"contactList"`


    // Clauses - Groups of conditions to filter the contacts by.
    Clauses []Contactlistfilterclause `json:"clauses"`


    // FilterType - How to join clauses together.
    FilterType string `json:"filterType"`


    

}

// String returns a JSON representation of the model
func (o *Contactlistfilter) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Clauses = []Contactlistfilterclause{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistfilter) MarshalJSON() ([]byte, error) {
    type Alias Contactlistfilter

    if ContactlistfilterMarshalled {
        return []byte("{}"), nil
    }
    ContactlistfilterMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        Version int `json:"version"`
        
        ContactList Domainentityref `json:"contactList"`
        
        Clauses []Contactlistfilterclause `json:"clauses"`
        
        FilterType string `json:"filterType"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Clauses: []Contactlistfilterclause{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


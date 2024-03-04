package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactbulksearchcriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactbulksearchcriteriaDud struct { 
    


    

}

// Contactbulksearchcriteria
type Contactbulksearchcriteria struct { 
    // Clauses - Groups of conditions to filter the contacts by.
    Clauses []Contactlistfilterclause `json:"clauses"`


    // FilterType - How to join clauses together.
    FilterType string `json:"filterType"`

}

// String returns a JSON representation of the model
func (o *Contactbulksearchcriteria) String() string {
     o.Clauses = []Contactlistfilterclause{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactbulksearchcriteria) MarshalJSON() ([]byte, error) {
    type Alias Contactbulksearchcriteria

    if ContactbulksearchcriteriaMarshalled {
        return []byte("{}"), nil
    }
    ContactbulksearchcriteriaMarshalled = true

    return json.Marshal(&struct {
        
        Clauses []Contactlistfilterclause `json:"clauses"`
        
        FilterType string `json:"filterType"`
        *Alias
    }{

        
        Clauses: []Contactlistfilterclause{{}},
        


        

        Alias: (*Alias)(u),
    })
}


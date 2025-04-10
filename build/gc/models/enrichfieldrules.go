package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EnrichfieldrulesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EnrichfieldrulesDud struct { 
    


    


    

}

// Enrichfieldrules
type Enrichfieldrules struct { 
    // DefaultAction - Default behavior for combining data from the submitted request with any entity found in the database. The default behavior if unspecified is `PreferProvided`, meaning any non-null fields in the submitted request will override data in the database, but all null fields will remain unchanged. Omitting a field in the request payload means that it will be treated as null.
    DefaultAction string `json:"defaultAction"`


    // Rules - Field-specific behaviors for how to combine data from different sources. For example, you can set a `defaultAction` of `PreferProvided`, but use different behaviors such as `AlwaysUseProvided` or `PreferExisting` for specific fields.
    Rules []Enrichfieldrule `json:"rules"`


    // DefaultArrayAction - Default behavior for combining items in array field from the submitted request with any array entity found in the database. The default behavior if unspecified is `fill`, meaning the field value will always be the partial concatenation of both the array in the Database and the array in the contact body, up to the size limit of the array
    DefaultArrayAction string `json:"defaultArrayAction"`

}

// String returns a JSON representation of the model
func (o *Enrichfieldrules) String() string {
    
     o.Rules = []Enrichfieldrule{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Enrichfieldrules) MarshalJSON() ([]byte, error) {
    type Alias Enrichfieldrules

    if EnrichfieldrulesMarshalled {
        return []byte("{}"), nil
    }
    EnrichfieldrulesMarshalled = true

    return json.Marshal(&struct {
        
        DefaultAction string `json:"defaultAction"`
        
        Rules []Enrichfieldrule `json:"rules"`
        
        DefaultArrayAction string `json:"defaultArrayAction"`
        *Alias
    }{

        


        
        Rules: []Enrichfieldrule{{}},
        


        

        Alias: (*Alias)(u),
    })
}


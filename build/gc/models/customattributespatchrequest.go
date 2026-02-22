package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomattributespatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomattributespatchrequestDud struct { 
    


    


    


    

}

// Customattributespatchrequest
type Customattributespatchrequest struct { 
    // Id - Unique identifier for the Custom Attributes record. IDs are created by users.
    Id string `json:"id"`


    // Divisions - The list of division ids. Use [] if divisions aren't used (Unassigned Division). Omitting or setting to [] clears existing values on update.
    Divisions []string `json:"divisions"`


    // Version - The latest version of the Custom Attributes record. Optional for concurrency check.
    Version int `json:"version"`


    // CustomAttributes - The map of attribute values.
    CustomAttributes map[string]interface{} `json:"customAttributes"`

}

// String returns a JSON representation of the model
func (o *Customattributespatchrequest) String() string {
    
     o.Divisions = []string{""} 
    
     o.CustomAttributes = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customattributespatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Customattributespatchrequest

    if CustomattributespatchrequestMarshalled {
        return []byte("{}"), nil
    }
    CustomattributespatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Divisions []string `json:"divisions"`
        
        Version int `json:"version"`
        
        CustomAttributes map[string]interface{} `json:"customAttributes"`
        *Alias
    }{

        


        
        Divisions: []string{""},
        


        


        
        CustomAttributes: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}


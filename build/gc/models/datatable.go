package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DatatableMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DatatableDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Datatable - Contains a metadata representation for a JSON schema stored in DataTables along with an optional field for the schema itself
type Datatable struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - The description from the JSON schema (equates to the Description field on the JSON schema.)
    Description string `json:"description"`


    // Schema - the schema as stored in the system.
    Schema Jsonschemadocument `json:"schema"`


    

}

// String returns a JSON representation of the model
func (o *Datatable) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Datatable) MarshalJSON() ([]byte, error) {
    type Alias Datatable

    if DatatableMarshalled {
        return []byte("{}"), nil
    }
    DatatableMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        Schema Jsonschemadocument `json:"schema"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


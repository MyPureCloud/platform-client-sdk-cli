package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchemaquantitylimitsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchemaquantitylimitsDud struct { 
    Id string `json:"id"`


    


    MinFieldNameCharacters int `json:"minFieldNameCharacters"`


    MaxFieldNameCharacters int `json:"maxFieldNameCharacters"`


    MinFieldDescriptionCharacters int `json:"minFieldDescriptionCharacters"`


    MaxFieldDescriptionCharacters int `json:"maxFieldDescriptionCharacters"`


    MinSchemaNameCharacters int `json:"minSchemaNameCharacters"`


    MaxSchemaNameCharacters int `json:"maxSchemaNameCharacters"`


    MinSchemaDescriptionCharacters int `json:"minSchemaDescriptionCharacters"`


    MaxSchemaDescriptionCharacters int `json:"maxSchemaDescriptionCharacters"`


    MaxNumberOfSchemasPerOrg int `json:"maxNumberOfSchemasPerOrg"`


    MaxNumberOfFieldsPerSchema int `json:"maxNumberOfFieldsPerSchema"`


    MaxNumberOfFieldsPerOrg int `json:"maxNumberOfFieldsPerOrg"`


    SelfUri string `json:"selfUri"`

}

// Schemaquantitylimits
type Schemaquantitylimits struct { 
    


    // Name
    Name string `json:"name"`


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Schemaquantitylimits) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schemaquantitylimits) MarshalJSON() ([]byte, error) {
    type Alias Schemaquantitylimits

    if SchemaquantitylimitsMarshalled {
        return []byte("{}"), nil
    }
    SchemaquantitylimitsMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


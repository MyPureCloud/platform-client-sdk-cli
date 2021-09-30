package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2schemaattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2schemaattributeDud struct { 
    Name string `json:"name"`


    VarType string `json:"type"`


    SubAttributes []Scimv2schemaattribute `json:"subAttributes"`


    MultiValued bool `json:"multiValued"`


    Description string `json:"description"`


    Required bool `json:"required"`


    CanonicalValues []string `json:"canonicalValues"`


    CaseExact bool `json:"caseExact"`


    Mutability string `json:"mutability"`


    Returned string `json:"returned"`


    Uniqueness string `json:"uniqueness"`


    ReferenceTypes []string `json:"referenceTypes"`

}

// Scimv2schemaattribute - A complex type that defines service provider attributes or subattributes and their qualities.
type Scimv2schemaattribute struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimv2schemaattribute) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2schemaattribute) MarshalJSON() ([]byte, error) {
    type Alias Scimv2schemaattribute

    if Scimv2schemaattributeMarshalled {
        return []byte("{}"), nil
    }
    Scimv2schemaattributeMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


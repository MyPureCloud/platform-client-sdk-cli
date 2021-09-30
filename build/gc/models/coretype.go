package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoretypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoretypeDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Coretype
type Coretype struct { 
    


    // Name
    Name string `json:"name"`


    // Version - A positive integer denoting the core type's version
    Version int `json:"version"`


    // DateCreated - The date the core type was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // Schema - The core type's built-in schema
    Schema Schema `json:"schema"`


    // Current - A boolean indicating if the core type's version is the current one in use by the system
    Current bool `json:"current"`


    // ValidationFields - An array of strings naming the fields of the core type subject to validation.  Validation constraints are specified by a schema author using the core type.
    ValidationFields []string `json:"validationFields"`


    // ValidationLimits - A structure denoting the system-imposed minimum and maximum string length (for text-based core types) or numeric values (for number-based) core types.  For example, the validationLimits for a text-based core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schemaauthor on a text field.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field.
    ValidationLimits Validationlimits `json:"validationLimits"`


    // ItemValidationFields - Specific to the \"tag\" core type, this is an array of strings naming the tag item fields of the core type subject to validation
    ItemValidationFields []string `json:"itemValidationFields"`


    // ItemValidationLimits - A structure denoting the system-imposed minimum and maximum string length for string-array based core types such as \"tag\" and \"enum\".  Forexample, the validationLimits for a schema field using a tag core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schema author on individual tags.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field's tags.
    ItemValidationLimits Itemvalidationlimits `json:"itemValidationLimits"`


    

}

// String returns a JSON representation of the model
func (o *Coretype) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ValidationFields = []string{""} 
    
    
    
    
    
    
    
     o.ItemValidationFields = []string{""} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coretype) MarshalJSON() ([]byte, error) {
    type Alias Coretype

    if CoretypeMarshalled {
        return []byte("{}"), nil
    }
    CoretypeMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        Schema Schema `json:"schema"`
        
        Current bool `json:"current"`
        
        ValidationFields []string `json:"validationFields"`
        
        ValidationLimits Validationlimits `json:"validationLimits"`
        
        ItemValidationFields []string `json:"itemValidationFields"`
        
        ItemValidationLimits Itemvalidationlimits `json:"itemValidationLimits"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        ValidationFields: []string{""},
        

        

        

        

        
        ItemValidationFields: []string{""},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


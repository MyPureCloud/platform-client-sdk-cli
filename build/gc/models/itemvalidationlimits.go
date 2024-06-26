package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ItemvalidationlimitsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ItemvalidationlimitsDud struct { 
    


    

}

// Itemvalidationlimits
type Itemvalidationlimits struct { 
    // MinLength - A structure denoting the system-imposed minimum string length (for text-based core types) or numeric values (for number-based) core types.  For example, the validationLimits for a text-based core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schemaauthor on a text field.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field.
    MinLength Minlength `json:"minLength"`


    // MaxLength - A structure denoting the system-imposed minimum and maximum string length (for text-based core types) or numeric values (for number-based) core types.  For example, the validationLimits for a text-based core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schemaauthor on a text field.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field.
    MaxLength Maxlength `json:"maxLength"`

}

// String returns a JSON representation of the model
func (o *Itemvalidationlimits) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Itemvalidationlimits) MarshalJSON() ([]byte, error) {
    type Alias Itemvalidationlimits

    if ItemvalidationlimitsMarshalled {
        return []byte("{}"), nil
    }
    ItemvalidationlimitsMarshalled = true

    return json.Marshal(&struct {
        
        MinLength Minlength `json:"minLength"`
        
        MaxLength Maxlength `json:"maxLength"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}


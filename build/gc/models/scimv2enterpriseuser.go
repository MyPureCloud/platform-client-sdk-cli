package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2enterpriseuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2enterpriseuserDud struct { 
    


    


    


    

}

// Scimv2enterpriseuser - Defines a SCIM enterprise user.
type Scimv2enterpriseuser struct { 
    // Division - The division that the user belongs to.
    Division string `json:"division"`


    // Department - The department that the user belongs to.
    Department string `json:"department"`


    // Manager - The user's manager.
    Manager Manager `json:"manager"`


    // EmployeeNumber - The user's employee number.
    EmployeeNumber string `json:"employeeNumber"`

}

// String returns a JSON representation of the model
func (o *Scimv2enterpriseuser) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2enterpriseuser) MarshalJSON() ([]byte, error) {
    type Alias Scimv2enterpriseuser

    if Scimv2enterpriseuserMarshalled {
        return []byte("{}"), nil
    }
    Scimv2enterpriseuserMarshalled = true

    return json.Marshal(&struct {
        
        Division string `json:"division"`
        
        Department string `json:"department"`
        
        Manager Manager `json:"manager"`
        
        EmployeeNumber string `json:"employeeNumber"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


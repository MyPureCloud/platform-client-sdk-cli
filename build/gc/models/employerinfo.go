package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmployerinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmployerinfoDud struct { 
    


    


    


    

}

// Employerinfo
type Employerinfo struct { 
    // OfficialName
    OfficialName string `json:"officialName"`


    // EmployeeId
    EmployeeId string `json:"employeeId"`


    // EmployeeType
    EmployeeType string `json:"employeeType"`


    // DateHire
    DateHire string `json:"dateHire"`

}

// String returns a JSON representation of the model
func (o *Employerinfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Employerinfo) MarshalJSON() ([]byte, error) {
    type Alias Employerinfo

    if EmployerinfoMarshalled {
        return []byte("{}"), nil
    }
    EmployerinfoMarshalled = true

    return json.Marshal(&struct { 
        OfficialName string `json:"officialName"`
        
        EmployeeId string `json:"employeeId"`
        
        EmployeeType string `json:"employeeType"`
        
        DateHire string `json:"dateHire"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


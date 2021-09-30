package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutofofficeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutofofficeDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Outofoffice
type Outofoffice struct { 
    


    // Name
    Name string `json:"name"`


    // User
    User *User `json:"user"`


    // StartDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // Active
    Active bool `json:"active"`


    // Indefinite
    Indefinite bool `json:"indefinite"`


    

}

// String returns a JSON representation of the model
func (o *Outofoffice) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outofoffice) MarshalJSON() ([]byte, error) {
    type Alias Outofoffice

    if OutofofficeMarshalled {
        return []byte("{}"), nil
    }
    OutofofficeMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        User *User `json:"user"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        Active bool `json:"active"`
        
        Indefinite bool `json:"indefinite"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


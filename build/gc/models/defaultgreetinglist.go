package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DefaultgreetinglistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DefaultgreetinglistDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Defaultgreetinglist
type Defaultgreetinglist struct { 
    


    // Name
    Name string `json:"name"`


    // Owner
    Owner Greetingowner `json:"owner"`


    // OwnerType
    OwnerType string `json:"ownerType"`


    // Greetings
    Greetings map[string]Greeting `json:"greetings"`


    // CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // CreatedBy
    CreatedBy string `json:"createdBy"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // ModifiedBy
    ModifiedBy string `json:"modifiedBy"`


    

}

// String returns a JSON representation of the model
func (o *Defaultgreetinglist) String() string {
    
    
    
     o.Greetings = map[string]Greeting{"": {}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Defaultgreetinglist) MarshalJSON() ([]byte, error) {
    type Alias Defaultgreetinglist

    if DefaultgreetinglistMarshalled {
        return []byte("{}"), nil
    }
    DefaultgreetinglistMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Owner Greetingowner `json:"owner"`
        
        OwnerType string `json:"ownerType"`
        
        Greetings map[string]Greeting `json:"greetings"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        CreatedBy string `json:"createdBy"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        ModifiedBy string `json:"modifiedBy"`
        *Alias
    }{

        


        


        


        


        
        Greetings: map[string]Greeting{"": {}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


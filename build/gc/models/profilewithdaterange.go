package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProfilewithdaterangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProfilewithdaterangeDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Profilewithdaterange
type Profilewithdaterange struct { 
    // Id - Profile ID
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    // DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEndWorkday time.Time `json:"dateEndWorkday"`


    

}

// String returns a JSON representation of the model
func (o *Profilewithdaterange) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Profilewithdaterange) MarshalJSON() ([]byte, error) {
    type Alias Profilewithdaterange

    if ProfilewithdaterangeMarshalled {
        return []byte("{}"), nil
    }
    ProfilewithdaterangeMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        DateStartWorkday time.Time `json:"dateStartWorkday"`
        
        DateEndWorkday time.Time `json:"dateEndWorkday"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


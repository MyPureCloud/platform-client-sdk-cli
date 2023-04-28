package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserprofilesindaterangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserprofilesindaterangeDud struct { 
    


    


    


    

}

// Userprofilesindaterange
type Userprofilesindaterange struct { 
    // User - The query agent
    User Userreference `json:"user"`


    // DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    // DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEndWorkday time.Time `json:"dateEndWorkday"`


    // Profiles - The list of profiles of the agent
    Profiles []Profilewithdaterange `json:"profiles"`

}

// String returns a JSON representation of the model
func (o *Userprofilesindaterange) String() string {
    
    
    
     o.Profiles = []Profilewithdaterange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userprofilesindaterange) MarshalJSON() ([]byte, error) {
    type Alias Userprofilesindaterange

    if UserprofilesindaterangeMarshalled {
        return []byte("{}"), nil
    }
    UserprofilesindaterangeMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        DateStartWorkday time.Time `json:"dateStartWorkday"`
        
        DateEndWorkday time.Time `json:"dateEndWorkday"`
        
        Profiles []Profilewithdaterange `json:"profiles"`
        *Alias
    }{

        


        


        


        
        Profiles: []Profilewithdaterange{{}},
        

        Alias: (*Alias)(u),
    })
}


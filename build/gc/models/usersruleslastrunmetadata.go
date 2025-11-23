package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersruleslastrunmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersruleslastrunmetadataDud struct { 
    


    


    


    

}

// Usersruleslastrunmetadata
type Usersruleslastrunmetadata struct { 
    // Id - The id of the run
    Id string `json:"id"`


    // CreatedDate - The date/time the rule was run. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // CreatedBy - The user that started the run
    CreatedBy Userreference `json:"createdBy"`


    // ResultCount - The number of users returned by the run
    ResultCount int `json:"resultCount"`

}

// String returns a JSON representation of the model
func (o *Usersruleslastrunmetadata) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersruleslastrunmetadata) MarshalJSON() ([]byte, error) {
    type Alias Usersruleslastrunmetadata

    if UsersruleslastrunmetadataMarshalled {
        return []byte("{}"), nil
    }
    UsersruleslastrunmetadataMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        ResultCount int `json:"resultCount"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidmetadataDud struct { 
    


    


    


    

}

// Workplanbidmetadata
type Workplanbidmetadata struct { 
    // CreatedBy - The user who created the associated entity
    CreatedBy Userreference `json:"createdBy"`


    // CreatedDate - The date the entity created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // ModifiedBy - The user who modified the entity
    ModifiedBy Userreference `json:"modifiedBy"`


    // ModifiedDate - The entity last modified date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`

}

// String returns a JSON representation of the model
func (o *Workplanbidmetadata) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidmetadata) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidmetadata

    if WorkplanbidmetadataMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidmetadataMarshalled = true

    return json.Marshal(&struct {
        
        CreatedBy Userreference `json:"createdBy"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}


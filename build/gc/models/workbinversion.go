package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkbinversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkbinversionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workbinversion
type Workbinversion struct { 
    


    // Name - Workbin name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - Workbin description
    Description string `json:"description"`


    // DateCreated - The creation date of the Workbin. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The modified date of the Workbin. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The id of the User who modified the Workbin.
    ModifiedBy Userreference `json:"modifiedBy"`


    // Version - Version
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Workbinversion) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workbinversion) MarshalJSON() ([]byte, error) {
    type Alias Workbinversion

    if WorkbinversionMarshalled {
        return []byte("{}"), nil
    }
    WorkbinversionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExecutiondataflowsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExecutiondataflowsettingsresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Executiondataflowsettingsresponse
type Executiondataflowsettingsresponse struct { 
    


    // Name
    Name string `json:"name"`


    // Enabled - whether or not the setting is enabled.
    Enabled bool `json:"enabled"`


    // ModifiedBy - User that last changed the setting.
    ModifiedBy Userreference `json:"modifiedBy"`


    // ModifiedByClient - OAuth client that last changed the setting.
    ModifiedByClient Domainentityref `json:"modifiedByClient"`


    // DateModified - The time this setting was set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Executiondataflowsettingsresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Executiondataflowsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Executiondataflowsettingsresponse

    if ExecutiondataflowsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    ExecutiondataflowsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        ModifiedByClient Domainentityref `json:"modifiedByClient"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


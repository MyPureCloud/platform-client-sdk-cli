package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicenowsourceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicenowsourceresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Servicenowsourceresponse
type Servicenowsourceresponse struct { 
    


    // Name - Name of the source.
    Name string `json:"name"`


    // DateCreated - Source creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Source last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // VarType - The source type.
    VarType string `json:"type"`


    // Integration - The reference to the integration associated with the source.
    Integration Knowledgeintegrationreference `json:"integration"`


    // SchedulePeriod - The schedule period of the source in hours.
    SchedulePeriod int `json:"schedulePeriod"`


    // LastSync - Additional information about the last synchronization of the source.
    LastSync Sourcelastsync `json:"lastSync"`


    // Settings - The settings of the source.
    Settings Servicenowsettings `json:"settings"`


    

}

// String returns a JSON representation of the model
func (o *Servicenowsourceresponse) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicenowsourceresponse) MarshalJSON() ([]byte, error) {
    type Alias Servicenowsourceresponse

    if ServicenowsourceresponseMarshalled {
        return []byte("{}"), nil
    }
    ServicenowsourceresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        VarType string `json:"type"`
        
        Integration Knowledgeintegrationreference `json:"integration"`
        
        SchedulePeriod int `json:"schedulePeriod"`
        
        LastSync Sourcelastsync `json:"lastSync"`
        
        Settings Servicenowsettings `json:"settings"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


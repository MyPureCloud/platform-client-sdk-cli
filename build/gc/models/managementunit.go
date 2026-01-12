package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManagementunitMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManagementunitDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    ModifiedBy Userreference `json:"modifiedBy"`


    Version int `json:"version"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Managementunit
type Managementunit struct { 
    


    // Name
    Name string `json:"name"`


    // BusinessUnit - The business unit to which this management unit belongs
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // StartDayOfWeek - Start day of week for scheduling and forecasting purposes. Moving to Business Unit
    StartDayOfWeek string `json:"startDayOfWeek"`


    // TimeZone - The time zone for the management unit in standard Olson format.  Moving to Business Unit
    TimeZone string `json:"timeZone"`


    // Settings - The configuration settings for this management unit
    Settings Managementunitsettingsresponse `json:"settings"`


    // Metadata - Version info metadata for this management unit. Deprecated, use settings.metadata
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // Division - The division to which this entity belongs.
    Division Divisionreference `json:"division"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Managementunit) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Managementunit) MarshalJSON() ([]byte, error) {
    type Alias Managementunit

    if ManagementunitMarshalled {
        return []byte("{}"), nil
    }
    ManagementunitMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        StartDayOfWeek string `json:"startDayOfWeek"`
        
        TimeZone string `json:"timeZone"`
        
        Settings Managementunitsettingsresponse `json:"settings"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        Division Divisionreference `json:"division"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


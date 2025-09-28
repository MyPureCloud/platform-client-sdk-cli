package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanlistitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanlistitemDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanlistitem
type Activityplanlistitem struct { 
    


    // Name - The name of the activity plan
    Name string `json:"name"`


    // ManagementUnits - The management units to which this activity plan applies. Empty list or null means this activity plan applies to all management units in the business unit
    ManagementUnits []Managementunitreference `json:"managementUnits"`


    // Description - The description of this activity plan
    Description string `json:"description"`


    // ActivityCode - The activity code to which this activity plan applies. Note: It is recommended to load and cache the entire list of activity codes rather than look up individual codes
    ActivityCode Activitycodereference `json:"activityCode"`


    // VarType - The type of the activity plan
    VarType string `json:"type"`


    // OptimizationObjective - The optimization objective of this activity plan
    OptimizationObjective string `json:"optimizationObjective"`


    // RecurrenceSettings - Recurrence settings for this activity plan
    RecurrenceSettings Recurrencesettings `json:"recurrenceSettings"`


    // State - The state of this activity plan
    State string `json:"state"`


    // LastRunDate - The date on which the activity plan was last manually run, in ISO-8601 format
    LastRunDate time.Time `json:"lastRunDate"`


    // LastRunBy - The last user to run this activity plan
    LastRunBy Userreference `json:"lastRunBy"`


    // CreatedDate - The date the activity plan was created, in ISO-8601 format
    CreatedDate time.Time `json:"createdDate"`


    // CreatedBy - The user who created this activity plan
    CreatedBy Userreference `json:"createdBy"`


    // ModifiedDate - The date the activity plan was modified, in ISO-8601 format
    ModifiedDate time.Time `json:"modifiedDate"`


    // ModifiedBy - The last user to modify this activity plan. The id may be 'System' if it was an automated process
    ModifiedBy Userreference `json:"modifiedBy"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanlistitem) String() string {
    
     o.ManagementUnits = []Managementunitreference{{}} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanlistitem) MarshalJSON() ([]byte, error) {
    type Alias Activityplanlistitem

    if ActivityplanlistitemMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanlistitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnits []Managementunitreference `json:"managementUnits"`
        
        Description string `json:"description"`
        
        ActivityCode Activitycodereference `json:"activityCode"`
        
        VarType string `json:"type"`
        
        OptimizationObjective string `json:"optimizationObjective"`
        
        RecurrenceSettings Recurrencesettings `json:"recurrenceSettings"`
        
        State string `json:"state"`
        
        LastRunDate time.Time `json:"lastRunDate"`
        
        LastRunBy Userreference `json:"lastRunBy"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        *Alias
    }{

        


        


        
        ManagementUnits: []Managementunitreference{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserschedulefulldaytimeoffmarkerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserschedulefulldaytimeoffmarkerDud struct { 
    


    


    


    


    


    

}

// Userschedulefulldaytimeoffmarker - Marker to indicate an approved full day time off request
type Userschedulefulldaytimeoffmarker struct { 
    // ManagementUnitDate - The date associated with the time off request that this marker corresponds to.  Date only, in ISO-8601 format.
    ManagementUnitDate string `json:"managementUnitDate"`


    // ActivityCodeId - The id for the activity code.  Look up a map of activity codes with the activities route
    ActivityCodeId string `json:"activityCodeId"`


    // IsPaid - Whether this is paid time off
    IsPaid bool `json:"isPaid"`


    // LengthInMinutes - The length in minutes of this time off marker
    LengthInMinutes int `json:"lengthInMinutes"`


    // Description - The description associated with the time off request that this marker corresponds to
    Description string `json:"description"`


    // Delete - If marked true for updating an existing full day time off marker, it will be deleted
    Delete bool `json:"delete"`

}

// String returns a JSON representation of the model
func (o *Userschedulefulldaytimeoffmarker) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userschedulefulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
    type Alias Userschedulefulldaytimeoffmarker

    if UserschedulefulldaytimeoffmarkerMarshalled {
        return []byte("{}"), nil
    }
    UserschedulefulldaytimeoffmarkerMarshalled = true

    return json.Marshal(&struct { 
        ManagementUnitDate string `json:"managementUnitDate"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        IsPaid bool `json:"isPaid"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        Description string `json:"description"`
        
        Delete bool `json:"delete"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


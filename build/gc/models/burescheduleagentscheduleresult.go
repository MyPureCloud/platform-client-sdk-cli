package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BurescheduleagentscheduleresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BurescheduleagentscheduleresultDud struct { 
    


    


    

}

// Burescheduleagentscheduleresult
type Burescheduleagentscheduleresult struct { 
    // ManagementUnit - The management unit to which this part of the result applies
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // DownloadResult - The agent schedules.  Result will always come via the downloadUrl; however the schema is included for documentation
    DownloadResult Murescheduleresultwrapper `json:"downloadResult"`


    // DownloadUrl - The download URL from which to fetch the result
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Burescheduleagentscheduleresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Burescheduleagentscheduleresult) MarshalJSON() ([]byte, error) {
    type Alias Burescheduleagentscheduleresult

    if BurescheduleagentscheduleresultMarshalled {
        return []byte("{}"), nil
    }
    BurescheduleagentscheduleresultMarshalled = true

    return json.Marshal(&struct { 
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        DownloadResult Murescheduleresultwrapper `json:"downloadResult"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}


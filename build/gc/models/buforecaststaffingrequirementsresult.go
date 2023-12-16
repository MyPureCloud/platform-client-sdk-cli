package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuforecaststaffingrequirementsresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuforecaststaffingrequirementsresultDud struct { 
    


    


    


    

}

// Buforecaststaffingrequirementsresult
type Buforecaststaffingrequirementsresult struct { 
    // WeekNumber - The week number represented by this response
    WeekNumber int `json:"weekNumber"`


    // DownloadUrl - The url to get the requirements results for this week
    DownloadUrl string `json:"downloadUrl"`


    // DownloadUrlExpirationDate - The expiration date of the download url, as an ISO-8601 string
    DownloadUrlExpirationDate time.Time `json:"downloadUrlExpirationDate"`


    // PlanningGroupStaffingRequirements - Results will always come via downloadUrl, however the schema is included for documentation
    PlanningGroupStaffingRequirements []Staffingrequirementsplanninggroupdata `json:"planningGroupStaffingRequirements"`

}

// String returns a JSON representation of the model
func (o *Buforecaststaffingrequirementsresult) String() string {
    
    
    
     o.PlanningGroupStaffingRequirements = []Staffingrequirementsplanninggroupdata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buforecaststaffingrequirementsresult) MarshalJSON() ([]byte, error) {
    type Alias Buforecaststaffingrequirementsresult

    if BuforecaststaffingrequirementsresultMarshalled {
        return []byte("{}"), nil
    }
    BuforecaststaffingrequirementsresultMarshalled = true

    return json.Marshal(&struct {
        
        WeekNumber int `json:"weekNumber"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadUrlExpirationDate time.Time `json:"downloadUrlExpirationDate"`
        
        PlanningGroupStaffingRequirements []Staffingrequirementsplanninggroupdata `json:"planningGroupStaffingRequirements"`
        *Alias
    }{

        


        


        


        
        PlanningGroupStaffingRequirements: []Staffingrequirementsplanninggroupdata{{}},
        

        Alias: (*Alias)(u),
    })
}


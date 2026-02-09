package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InactivitytimeoutgroupbundleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InactivitytimeoutgroupbundleDud struct { 
    


    


    


    

}

// Inactivitytimeoutgroupbundle
type Inactivitytimeoutgroupbundle struct { 
    // Priority - The priority of the group bundle (1-5).
    Priority int `json:"priority"`


    // TimeoutSeconds - The timeout value in seconds (300 to 28800, representing 5 to 480 minutes).
    TimeoutSeconds int `json:"timeoutSeconds"`


    // InactivityTimeoutUnit - The unit for the timeout (MINUTES or HOURS).
    InactivityTimeoutUnit string `json:"inactivityTimeoutUnit"`


    // Groups - The list of group IDs to select.
    Groups []string `json:"groups"`

}

// String returns a JSON representation of the model
func (o *Inactivitytimeoutgroupbundle) String() string {
    
    
    
     o.Groups = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Inactivitytimeoutgroupbundle) MarshalJSON() ([]byte, error) {
    type Alias Inactivitytimeoutgroupbundle

    if InactivitytimeoutgroupbundleMarshalled {
        return []byte("{}"), nil
    }
    InactivitytimeoutgroupbundleMarshalled = true

    return json.Marshal(&struct {
        
        Priority int `json:"priority"`
        
        TimeoutSeconds int `json:"timeoutSeconds"`
        
        InactivityTimeoutUnit string `json:"inactivityTimeoutUnit"`
        
        Groups []string `json:"groups"`
        *Alias
    }{

        


        


        


        
        Groups: []string{""},
        

        Alias: (*Alias)(u),
    })
}


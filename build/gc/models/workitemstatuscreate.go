package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemstatuscreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemstatuscreateDud struct { 
    


    


    


    


    


    


    


    

}

// Workitemstatuscreate
type Workitemstatuscreate struct { 
    // Name - The name of the Status. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // Category - The Category of the Status.
    Category string `json:"category"`


    // DestinationStatusIds - A list of destination Statuses where a Workitem with this Status can transition to. If the list is empty Workitems with this Status can transition to all other Statuses defined on the Worktype. A Status can have a maximum of 24 destinations.
    DestinationStatusIds []string `json:"destinationStatusIds"`


    // Description - The description of the Status. Maximum length of 512 characters.
    Description string `json:"description"`


    // DefaultDestinationStatusId - Default destination status to which this Status will transition to if auto status transition enabled.
    DefaultDestinationStatusId string `json:"defaultDestinationStatusId"`


    // StatusTransitionDelaySeconds - Delay in seconds for auto status transition. Required if defaultDestinationStatusId is provided.
    StatusTransitionDelaySeconds int `json:"statusTransitionDelaySeconds"`


    // StatusTransitionTime - Time is represented as an ISO-8601 string without a timezone. For example: HH:mm:ss.SSS
    StatusTransitionTime string `json:"statusTransitionTime"`


    // AutoTerminateWorkitem - Terminate workitem on selection of status. Applicable only for statuses in the Closed category.
    AutoTerminateWorkitem bool `json:"autoTerminateWorkitem"`

}

// String returns a JSON representation of the model
func (o *Workitemstatuscreate) String() string {
    
    
     o.DestinationStatusIds = []string{""} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemstatuscreate) MarshalJSON() ([]byte, error) {
    type Alias Workitemstatuscreate

    if WorkitemstatuscreateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemstatuscreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Category string `json:"category"`
        
        DestinationStatusIds []string `json:"destinationStatusIds"`
        
        Description string `json:"description"`
        
        DefaultDestinationStatusId string `json:"defaultDestinationStatusId"`
        
        StatusTransitionDelaySeconds int `json:"statusTransitionDelaySeconds"`
        
        StatusTransitionTime string `json:"statusTransitionTime"`
        
        AutoTerminateWorkitem bool `json:"autoTerminateWorkitem"`
        *Alias
    }{

        


        


        
        DestinationStatusIds: []string{""},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


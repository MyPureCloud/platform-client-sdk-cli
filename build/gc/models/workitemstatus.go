package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemstatusDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workitemstatus
type Workitemstatus struct { 
    


    // Name
    Name string `json:"name"`


    // Category - The Category of the Status.
    Category string `json:"category"`


    // DestinationStatuses - The Statuses the Status can transition to.
    DestinationStatuses []Workitemstatusreference `json:"destinationStatuses"`


    // Description - The description of the Status.
    Description string `json:"description"`


    // DefaultDestinationStatus - Default destination status to which this Status will transition to if auto status transition enabled.
    DefaultDestinationStatus Workitemstatusreference `json:"defaultDestinationStatus"`


    // StatusTransitionDelaySeconds - Delay in seconds for auto status transition
    StatusTransitionDelaySeconds int `json:"statusTransitionDelaySeconds"`


    // StatusTransitionTime - Time is represented as an ISO-8601 string without a timezone. For example: HH:mm:ss.SSS
    StatusTransitionTime string `json:"statusTransitionTime"`


    // Worktype - The Worktype containing the Status.
    Worktype Worktypereference `json:"worktype"`


    // AutoTerminateWorkitem - Terminate workitem on selection of status. Applicable only for statuses in the Closed category.
    AutoTerminateWorkitem bool `json:"autoTerminateWorkitem"`


    

}

// String returns a JSON representation of the model
func (o *Workitemstatus) String() string {
    
    
     o.DestinationStatuses = []Workitemstatusreference{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemstatus) MarshalJSON() ([]byte, error) {
    type Alias Workitemstatus

    if WorkitemstatusMarshalled {
        return []byte("{}"), nil
    }
    WorkitemstatusMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Category string `json:"category"`
        
        DestinationStatuses []Workitemstatusreference `json:"destinationStatuses"`
        
        Description string `json:"description"`
        
        DefaultDestinationStatus Workitemstatusreference `json:"defaultDestinationStatus"`
        
        StatusTransitionDelaySeconds int `json:"statusTransitionDelaySeconds"`
        
        StatusTransitionTime string `json:"statusTransitionTime"`
        
        Worktype Worktypereference `json:"worktype"`
        
        AutoTerminateWorkitem bool `json:"autoTerminateWorkitem"`
        *Alias
    }{

        


        


        


        
        DestinationStatuses: []Workitemstatusreference{{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


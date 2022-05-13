package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatroutingtargetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatroutingtargetDud struct { 
    


    


    


    


    

}

// Webchatroutingtarget
type Webchatroutingtarget struct { 
    // TargetType - The target type of the routing target, such as 'QUEUE'.
    TargetType string `json:"targetType"`


    // TargetAddress - The target of the route, in the format appropriate given the 'targetType'.
    TargetAddress string `json:"targetAddress"`


    // Skills - The list of skill names to use for routing.
    Skills []string `json:"skills"`


    // Language - The language name to use for routing.
    Language string `json:"language"`


    // Priority - The priority to assign to the conversation for routing.
    Priority int `json:"priority"`

}

// String returns a JSON representation of the model
func (o *Webchatroutingtarget) String() string {
    
    
     o.Skills = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatroutingtarget) MarshalJSON() ([]byte, error) {
    type Alias Webchatroutingtarget

    if WebchatroutingtargetMarshalled {
        return []byte("{}"), nil
    }
    WebchatroutingtargetMarshalled = true

    return json.Marshal(&struct {
        
        TargetType string `json:"targetType"`
        
        TargetAddress string `json:"targetAddress"`
        
        Skills []string `json:"skills"`
        
        Language string `json:"language"`
        
        Priority int `json:"priority"`
        *Alias
    }{

        


        


        
        Skills: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutepathrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutepathrequestDud struct { 
    


    


    


    


    

}

// Routepathrequest
type Routepathrequest struct { 
    // QueueId - The ID of the queue to associate with the route path
    QueueId string `json:"queueId"`


    // MediaType - The media type of the given queue to associate with the route path
    MediaType string `json:"mediaType"`


    // LanguageId - The ID of the language to associate with the route path
    LanguageId string `json:"languageId"`


    // SkillIds - The set of skill IDs to associate with the route path
    SkillIds []string `json:"skillIds"`


    // SourcePlanningGroup - The planning group from which to take route paths. This property is only needed if a route path already exists in another planning group.Note that taking a route path from another planning group will modify the other planning group
    SourcePlanningGroup Sourceplanninggrouprequest `json:"sourcePlanningGroup"`

}

// String returns a JSON representation of the model
func (o *Routepathrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.SkillIds = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routepathrequest) MarshalJSON() ([]byte, error) {
    type Alias Routepathrequest

    if RoutepathrequestMarshalled {
        return []byte("{}"), nil
    }
    RoutepathrequestMarshalled = true

    return json.Marshal(&struct { 
        QueueId string `json:"queueId"`
        
        MediaType string `json:"mediaType"`
        
        LanguageId string `json:"languageId"`
        
        SkillIds []string `json:"skillIds"`
        
        SourcePlanningGroup Sourceplanninggrouprequest `json:"sourcePlanningGroup"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        SkillIds: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}


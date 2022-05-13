package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutepathresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutepathresponseDud struct { 
    


    


    


    

}

// Routepathresponse
type Routepathresponse struct { 
    // Queue - The ID of the queue associated with the route path
    Queue Queuereference `json:"queue"`


    // MediaType - The media type of the given queue associated with the route path
    MediaType string `json:"mediaType"`


    // Language - The ID of the language associated with the route path
    Language Languagereference `json:"language"`


    // Skills - The set of skills associated with the route path
    Skills []Routingskillreference `json:"skills"`

}

// String returns a JSON representation of the model
func (o *Routepathresponse) String() string {
    
    
    
     o.Skills = []Routingskillreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routepathresponse) MarshalJSON() ([]byte, error) {
    type Alias Routepathresponse

    if RoutepathresponseMarshalled {
        return []byte("{}"), nil
    }
    RoutepathresponseMarshalled = true

    return json.Marshal(&struct {
        
        Queue Queuereference `json:"queue"`
        
        MediaType string `json:"mediaType"`
        
        Language Languagereference `json:"language"`
        
        Skills []Routingskillreference `json:"skills"`
        *Alias
    }{

        


        


        


        
        Skills: []Routingskillreference{{}},
        

        Alias: (*Alias)(u),
    })
}


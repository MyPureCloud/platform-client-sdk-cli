package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuideDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    Status string `json:"status"`


    Source string `json:"source"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`


    LatestSavedVersion Guideversionref `json:"latestSavedVersion"`


    LatestProductionReadyVersion Guideversionref `json:"latestProductionReadyVersion"`

}

// Guide
type Guide struct { 
    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Guide) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guide) MarshalJSON() ([]byte, error) {
    type Alias Guide

    if GuideMarshalled {
        return []byte("{}"), nil
    }
    GuideMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


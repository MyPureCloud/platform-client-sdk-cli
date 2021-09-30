package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentupdateDud struct { 
    


    


    


    


    


    


    


    


    

}

// Documentupdate
type Documentupdate struct { 
    // ChangeNumber
    ChangeNumber int `json:"changeNumber"`


    // Name - The name of the document
    Name string `json:"name"`


    // Read
    Read bool `json:"read"`


    // AddTags
    AddTags []string `json:"addTags"`


    // RemoveTags
    RemoveTags []string `json:"removeTags"`


    // AddTagIds
    AddTagIds []string `json:"addTagIds"`


    // RemoveTagIds
    RemoveTagIds []string `json:"removeTagIds"`


    // UpdateAttributes
    UpdateAttributes []Documentattribute `json:"updateAttributes"`


    // RemoveAttributes
    RemoveAttributes []string `json:"removeAttributes"`

}

// String returns a JSON representation of the model
func (o *Documentupdate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AddTags = []string{""} 
    
    
    
     o.RemoveTags = []string{""} 
    
    
    
     o.AddTagIds = []string{""} 
    
    
    
     o.RemoveTagIds = []string{""} 
    
    
    
     o.UpdateAttributes = []Documentattribute{{}} 
    
    
    
     o.RemoveAttributes = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentupdate) MarshalJSON() ([]byte, error) {
    type Alias Documentupdate

    if DocumentupdateMarshalled {
        return []byte("{}"), nil
    }
    DocumentupdateMarshalled = true

    return json.Marshal(&struct { 
        ChangeNumber int `json:"changeNumber"`
        
        Name string `json:"name"`
        
        Read bool `json:"read"`
        
        AddTags []string `json:"addTags"`
        
        RemoveTags []string `json:"removeTags"`
        
        AddTagIds []string `json:"addTagIds"`
        
        RemoveTagIds []string `json:"removeTagIds"`
        
        UpdateAttributes []Documentattribute `json:"updateAttributes"`
        
        RemoveAttributes []string `json:"removeAttributes"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        AddTags: []string{""},
        

        

        
        RemoveTags: []string{""},
        

        

        
        AddTagIds: []string{""},
        

        

        
        RemoveTagIds: []string{""},
        

        

        
        UpdateAttributes: []Documentattribute{{}},
        

        

        
        RemoveAttributes: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}


package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesyncjobstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesyncjobstatisticsDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Knowledgesyncjobstatistics
type Knowledgesyncjobstatistics struct { 
    // CountDocumentImportActivityCreate - Number of documents will be created by the sync.
    CountDocumentImportActivityCreate int `json:"countDocumentImportActivityCreate"`


    // CountDocumentImportActivityUpdate - Number of documents will be updated by the sync.
    CountDocumentImportActivityUpdate int `json:"countDocumentImportActivityUpdate"`


    // CountDocumentStateDraft - Number of documents will be imported as draft.
    CountDocumentStateDraft int `json:"countDocumentStateDraft"`


    // CountDocumentStatePublished - Number of documents will be imported as published.
    CountDocumentStatePublished int `json:"countDocumentStatePublished"`


    // CountDocumentImportSuccess - Number of imported documents.
    CountDocumentImportSuccess int `json:"countDocumentImportSuccess"`


    // CountDocumentImportFailure - Number of documents failed to import.
    CountDocumentImportFailure int `json:"countDocumentImportFailure"`


    // CountCategoryImportSuccess - Number of imported categories.
    CountCategoryImportSuccess int `json:"countCategoryImportSuccess"`


    // CountCategoryImportFailure - Number of categories failed to import.
    CountCategoryImportFailure int `json:"countCategoryImportFailure"`


    // CountLabelImportSuccess - Number of imported labels.
    CountLabelImportSuccess int `json:"countLabelImportSuccess"`


    // CountLabelImportFailure - Number of labels failed to import.
    CountLabelImportFailure int `json:"countLabelImportFailure"`


    // CountDocumentDeleteSuccess - Number of documents will be deleted by the sync.
    CountDocumentDeleteSuccess int `json:"countDocumentDeleteSuccess"`


    // CountDocumentDeleteFailure - Number of documents failed to delete.
    CountDocumentDeleteFailure int `json:"countDocumentDeleteFailure"`


    // CountCategoryDeleteSuccess - Number of successfully deleted categories.
    CountCategoryDeleteSuccess int `json:"countCategoryDeleteSuccess"`


    // CountCategoryDeleteFailure - Number of categories failed to delete.
    CountCategoryDeleteFailure int `json:"countCategoryDeleteFailure"`


    // CountLabelDeleteSuccess - Number of successfully deleted labels.
    CountLabelDeleteSuccess int `json:"countLabelDeleteSuccess"`


    // CountLabelDeleteFailure - Number of labels failed to delete.
    CountLabelDeleteFailure int `json:"countLabelDeleteFailure"`

}

// String returns a JSON representation of the model
func (o *Knowledgesyncjobstatistics) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesyncjobstatistics) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesyncjobstatistics

    if KnowledgesyncjobstatisticsMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesyncjobstatisticsMarshalled = true

    return json.Marshal(&struct {
        
        CountDocumentImportActivityCreate int `json:"countDocumentImportActivityCreate"`
        
        CountDocumentImportActivityUpdate int `json:"countDocumentImportActivityUpdate"`
        
        CountDocumentStateDraft int `json:"countDocumentStateDraft"`
        
        CountDocumentStatePublished int `json:"countDocumentStatePublished"`
        
        CountDocumentImportSuccess int `json:"countDocumentImportSuccess"`
        
        CountDocumentImportFailure int `json:"countDocumentImportFailure"`
        
        CountCategoryImportSuccess int `json:"countCategoryImportSuccess"`
        
        CountCategoryImportFailure int `json:"countCategoryImportFailure"`
        
        CountLabelImportSuccess int `json:"countLabelImportSuccess"`
        
        CountLabelImportFailure int `json:"countLabelImportFailure"`
        
        CountDocumentDeleteSuccess int `json:"countDocumentDeleteSuccess"`
        
        CountDocumentDeleteFailure int `json:"countDocumentDeleteFailure"`
        
        CountCategoryDeleteSuccess int `json:"countCategoryDeleteSuccess"`
        
        CountCategoryDeleteFailure int `json:"countCategoryDeleteFailure"`
        
        CountLabelDeleteSuccess int `json:"countLabelDeleteSuccess"`
        
        CountLabelDeleteFailure int `json:"countLabelDeleteFailure"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}


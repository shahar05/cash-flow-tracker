package core

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadAndValidate(w http.ResponseWriter, r *http.Request, modelRef interface{}) bool {
	body := r.Body
	defer body.Close()
	if err := json.NewDecoder(body).Decode(modelRef); err != nil {
		log.Printf("ReadAndValidate Error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	log.Printf("ReadAndValidate: %v", modelRef)

	// TODO: Validate Struct

	return true
}

// func ReadAndValidate(ctx iris.Context, modelRef interface{}) bool {

// 	defer r.Body.Close()
// 	b, _ := io.ReadAll(r.Body)
// 	json.Unmarshal([]byte(b), modelRef)

// 	if err := ctx.ReadJSON(modelRef); err != nil {
// 		golog.Debug(err)
// 		ctx.StatusCode(iris.StatusBadRequest)
// 		return false
// 	}

// 	golog.Debugf("Model: %+v", modelRef)

// 	if v := ValidateStruct(modelRef); v != nil {
// 		ctx.StatusCode(iris.StatusBadRequest)
// 		ctx.JSON(v)
// 		return false
// 	}

// 	return true
// }

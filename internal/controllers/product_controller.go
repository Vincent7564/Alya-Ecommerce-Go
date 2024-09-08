package controllers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"time"
// )

// func (uc *UserController) RegisterProduct(w http.ResponseWriter, r *http.Request) {
// 	var product struct {
// 		ProductName  string  `json:"product_name"`
// 		ProductPrice int8    `json:"product_price"`
// 		Discount     float32 `json:"discount"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&product)
// 	if err != nil {
// 		http.Error(w, "Invalid Request, Please Check", http.StatusBadRequest)
// 		return
// 	}

// 	t := time.Now()

// 	data,_,err :=  uc.Service.Client.From("products").Insert(interface{}{

// 	},false,"","","").Execute()

// }

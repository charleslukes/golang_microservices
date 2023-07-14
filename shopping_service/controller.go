package main

type payload struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := &Shop{}

	err := decoder.Decode(params)

	if err != nil {
		helper.RespondWithError(w, 400, err.Error())
		return
	}

	vErr := params.CreateOrderValidate()

	if vErr != nil {
		helper.RespondWithError(w, 400, vErr.Error())
		return
	}

	res, dbErr := mh.AddOne(params)

	if dbErr != nil {
		helper.RespondWithError(w, 404, dbErr.Error())
		return
	}

	payload := payload{
		Message: "order created",
		Data:    res,
	}

	helper.RespondWithJson(w, 201, payload)
}

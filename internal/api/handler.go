package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ajay-patidar-0/rag/internal/rag"
	"github.com/ajay-patidar-0/rag/internal/utils"
)

type QueryRequest struct {
	Query string `json:"query"`
}

func (as *ApiServer) QueryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle the query")

	var Resp QueryRequest
	if err := json.NewDecoder(r.Body).Decode(&Resp); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid format")
		return
	}

	if Resp.Query == "" {
		utils.RespondError(w, http.StatusBadRequest, "please send some query")
	}

	result, err := rag.QuerytoAnswer(Resp.Query, as.Db.DB)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Sorry!, there is a server problem")
		return
	}

	utils.RespondSuccess(w, http.StatusAccepted, result)
}

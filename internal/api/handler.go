package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	paperextractor "github.com/ajay-patidar-0/rag/internal/paper_extractor"
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
		return
	}

	result, err := rag.QuerytoAnswer(Resp.Query, as.Db.DB)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Sorry!, there is a server problem")
		return
	}

	utils.RespondSuccess(w, http.StatusAccepted, result)
}

func (as *ApiServer) ExamPaperHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	examPaper, _, err := r.FormFile("file")
	if err != nil {
		utils.RespondError(w, http.StatusNotAcceptable, "file not found")
		return
	}
	defer examPaper.Close()

	var base64string string
	paperBytes, err := io.ReadAll(examPaper)
	if err != nil {
		log.Println(err)
		utils.RespondError(w, http.StatusNotAcceptable, "inappropriate content")
		return
	}
	base64string = base64.StdEncoding.EncodeToString(paperBytes)

	extractedPaper, err := paperextractor.PaperImageToJson(base64string)
	if err != nil {
		log.Println(err)
		utils.RespondError(w, http.StatusInternalServerError, "their is server problem can't able to extract paper")
		return
	}
	err = paperextractor.StorePaper(extractedPaper, as.Db)
	if err != nil {
		log.Println(err)
		utils.RespondError(w, http.StatusInternalServerError, "paper can't be added in database. Server problem")
		return
	}

	utils.RespondSuccess(w, http.StatusOK, "paper stored in database")
}

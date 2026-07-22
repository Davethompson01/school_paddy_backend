package handler

import (
	"encoding/json"
	"net/http"
	"time"

	middleware "github.com/Davethompson01/School_Paddy_golang/Middleware"
	auth "github.com/Davethompson01/School_Paddy_golang/internal/Auth"
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	Services "github.com/Davethompson01/School_Paddy_golang/internal/services"
)

func Upload_homework(apiCfg *config.ApiConfig) http.HandlerFunc {
	return func(res http.ResponseWriter, r *http.Request) {
		var project students.Project
		if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		claims := r.Context().Value(middleware.ClaimsKey).(*auth.Claims)
		project.UserID = claims.UserID
		project.UpdatedAt = time.Now()
		uploadService, err := Services.Upload_homework(apiCfg, project)
		if err != nil {
			RespondWithJson(res, http.StatusUnauthorized, false, err.Error(), nil)
			return
		}
		RespondWithJson(res, 201, true, "Uploaded Homework", uploadService)
	}
}

func HandlerCreateBID(api *config.ApiConfig) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		var bid solutionexpert_model.ApplyForHomeWork
		if err := json.NewDecoder(req.Body).Decode(&bid); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		claims := req.Context().Value(middleware.ClaimsKey).(*auth.Claims)
		bid.Solution_expert_id = claims.UserID

		createBid, err := Services.Service_CreateHomeWorkBID(api, bid)
		if err != nil {
			RespondWithJson(res, http.StatusUnauthorized, false, err.Error(), nil)
			return
		}
		RespondWithJson(res, 201, true, "Applied for Homework", createBid)
	}
}

func HandlerNegotiateBID(api *config.ApiConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var bid solutionexpert_model.NegotiateProject

		if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		claims := r.Context().Value(middleware.ClaimsKey).(*auth.Claims)
		bid.Solution_expert_id = claims.UserID

		negotiateBid, err := Services.NegotiateBid(api, bid)
		if err != nil {
			RespondWithJson(w, http.StatusUnauthorized, false, err.Error(), nil)
			return
		}
		RespondWithJson(w, http.StatusAccepted, true, "Negotiated for Homework", negotiateBid)
	}
}

func HandlerAcceptBID(api *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bid students.AcceptBid_HomeWork

		if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		claims := r.Context().Value(middleware.ClaimsKey).(*auth.Claims)
		bid.Student_id = claims.UserID

		acceptBid, err := Services.AcceptBID(api, bid)
		if err != nil {
			RespondWithJson(w, http.StatusUnauthorized, false, err.Error(), nil)
			return
		}
		RespondWithJson(w, http.StatusAccepted, true, "Solution Expert accepted", acceptBid)

	}

}
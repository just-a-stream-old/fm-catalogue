package api

import (
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/model"
	"net/http"
)

func (s *server) getExchanges() http.HandlerFunc {
	type queryParams struct {

	}

	type response struct {
		Exchanges []model.Exchange
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		exchanges, _, err := s.fMService.GetExchanges(ctx)
		if err != nil {
			s.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

		response := response{
			Exchanges: exchanges,
		}

		s.respondJSON(w, http.StatusOK, response)
	}
}

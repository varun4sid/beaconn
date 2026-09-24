package leave

import (
	"fmt"
	"net/http"
	"time"

	"github.com/varun4sid/beaconn-backend/internal/api"
	cfg "github.com/varun4sid/beaconn-backend/internal/config"
	"github.com/varun4sid/beaconn-backend/internal/scraper"
)

func leaveHandler(w http.ResponseWriter, r *http.Request, c *cfg.Config) {
	var payload scraper.LeavePayload

	body, err := api.ParseRequestBody[LeaveRequest](r)
	if err != nil {
		api.HttpError(w, err, http.StatusBadRequest)
		return
	}

	if body.NumDays == nil && (body.FromDate == nil || body.ToDate == nil) {
		api.HttpError(w, fmt.Errorf("either num_days or both from_date and to_date must be provided"), http.StatusBadRequest)
		return
	}

	if body.NumDays != nil {
		if *body.NumDays <= 0 {
			api.HttpError(w, fmt.Errorf("num_days must be a positive integer"), http.StatusBadRequest)
			return
		} else if *body.NumDays > 30 {
			api.HttpError(w, fmt.Errorf("num_days must not exceed 30"), http.StatusBadRequest)
			return
		} else {
			IST := time.FixedZone("IST", 19800)
			now := time.Now().In(IST)
			payload.FromDate = now.Format("02/01/2006")
			payload.ToDate = now.AddDate(0, 0, *body.NumDays-1).Format("02/01/2006")
		}
	}

	if body.FromDate != nil {
		IST := time.FixedZone("IST", 19800)
		fromDate, err := time.ParseInLocation("02/01/2006", *body.FromDate, IST)
		if err != nil {
			api.HttpError(w, fmt.Errorf("invalid from_date format: %w", err), http.StatusBadRequest)
			return
		}

		toDate, err := time.ParseInLocation("02/01/2006", *body.ToDate, IST)
		if err != nil {
			api.HttpError(w, fmt.Errorf("invalid to_date format: %w", err), http.StatusBadRequest)
			return
		}

		if toDate.Before(fromDate) {
			api.HttpError(w, fmt.Errorf("to_date cannot be before from_date"), http.StatusBadRequest)
			return
		}

		payload.FromDate = fromDate.Format("02/01/2006")
		payload.ToDate = toDate.Format("02/01/2006")
	}

	s, err := scraper.NewSession(c)
	if err != nil {
		api.HttpError(w, fmt.Errorf("failed to create scraper session: %w", err), http.StatusInternalServerError)
		return
	}
	
	statusCode, err := s.Login(body.Rollno, body.Password)
	if err != nil || statusCode != http.StatusOK {
		api.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	if body.Reason == nil {
		payload.Reason = "Leave"
	} else {
		payload.Reason = *body.Reason
	}

	payload.Rollno = body.Rollno
	payload.FromTime = "8:00"
	payload.ToTime = "8:00"
	payload.FromSession = "AM"
	payload.ToSession = "AM"
	payload.LeaveType = "L"
	payload.Manager, err = s.GetManager()

	if err != nil {
		api.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	err = s.ApplyLeave(payload)
	if err != nil {
		api.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"status":  "success",
		"message": "Leave applied successfully",
	}

	api.WriteResponseBody(w, response)
}
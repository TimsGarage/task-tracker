package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

const maximumRequestBodySize = 1 << 20

func parsePathID(
	r *http.Request,
	name string,
) (uint, error) {
	value := r.PathValue(name)

	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("invalid ID")
	}

	return uint(id), nil
}

func decodeJSON(
	w http.ResponseWriter,
	r *http.Request,
	destination any,
) bool {
	r.Body = http.MaxBytesReader(
		w,
		r.Body,
		maximumRequestBodySize,
	)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(destination); err != nil {
		badRequest(w, "Invalid JSON")
		return false
	}

	var extra any

	err := decoder.Decode(&extra)
	if !errors.Is(err, io.EOF) {
		badRequest(w, "Request body must contain one JSON object")
		return false
	}

	return true
}

func writeJSON(
	w http.ResponseWriter,
	status int,
	value any,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if value == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func writeError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}

func badRequest(
	w http.ResponseWriter,
	message string,
) {
	writeError(w, http.StatusBadRequest, message)
}

func notFound(
	w http.ResponseWriter,
	message string,
) {
	writeError(w, http.StatusNotFound, message)
}

func serverError(
	w http.ResponseWriter,
	err error,
) {
	log.Printf("internal server error: %v", err)

	writeError(
		w,
		http.StatusInternalServerError,
		"Internal server error",
	)
}

func repositoryError(
	w http.ResponseWriter,
	err error,
	notFoundMessage string,
) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound(w, notFoundMessage)
		return
	}

	serverError(w, err)
}

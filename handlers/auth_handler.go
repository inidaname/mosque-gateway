package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/inidaname/mosque/api_gateway/pkg/types"
	"github.com/inidaname/mosque/api_gateway/pkg/utils"
	pb "github.com/inidaname/mosque/protos"
)

func RegisterUser(client pb.AuthServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.CreateUserPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := payload.Validate(); err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		resp, err := client.RegisterUser(ctx, &pb.RegisterUserRequest{
			Email:    payload.Email,
			Password: payload.Password,
			Phone:    &payload.Phone,
			FullName: payload.FullName,
		})

		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}
		// json.NewEncoder(w).Encode(resp)
		utils.SendResponse(w, http.StatusCreated, "User registered successfully", resp)
	}
}

func LoginUser(client pb.AuthServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.LoginUserPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := payload.Validate(); err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		resp, err := client.LoginUser(ctx, &pb.LoginUserRequest{
			Email:    payload.Email,
			Password: payload.Password,
		})

		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		utils.SendResponse(w, http.StatusCreated, "User loggedin successfully", resp)
	}
}

func ForgotPassword(client pb.AuthServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.ForgotPasswordPayload
		if err := payload.Validate(); err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		resp, err := client.ForgotPassword(r.Context(), &pb.ForgotPasswordRequest{Email: payload.Email})

		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		utils.SendResponse(w, http.StatusCreated, "User loggedin successfully", resp)
	}
}

func ValidatePasswordToken(client pb.AuthServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var payload types.ValidatePasswordToken
		token := chi.URLParam(r, "token")

		resp, err := client.ValidatePasswordToken(r.Context(), &pb.ValidatePasswordTokenRequest{Token: token})

		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		utils.SendResponse(w, http.StatusCreated, "User loggedin successfully", resp)
	}
}

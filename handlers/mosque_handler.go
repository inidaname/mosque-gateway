package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/inidaname/mosque/api_gateway/pkg/types"
	"github.com/inidaname/mosque/api_gateway/pkg/utils"
	pb "github.com/inidaname/mosque_location/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateMosque(client pb.MosqueServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.CreateMosquePayload

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

		resp, err := client.CreateMosque(ctx, &pb.CreateMosqueRequest{
			Name:       payload.Name,
			Address:    payload.Address,
			EidTime:    &timestamppb.Timestamp{Seconds: payload.EidTime.Unix()},
			JummahTime: &timestamppb.Timestamp{Seconds: payload.JummahTime.Unix()},
			Lat:        payload.Lat,
			Lng:        payload.Lng,
		})

		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}
		// json.NewEncoder(w).Encode(resp)
		utils.SendResponse(w, http.StatusCreated, "Mosque created successfully", resp)
	}
}

func ListMosque(client pb.MosqueServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		resp, err := client.ListMosques(ctx, &pb.ListMosquesRequest{})

		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}
		// json.NewEncoder(w).Encode(resp)
		utils.SendResponse(w, http.StatusCreated, "Fetched Mosque successfully", resp)
	}
}

func UpdateMosque(client pb.MosqueServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.CreateMosquePayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		mosqueId, err := utils.ReadIDParam(r, "mosqueId")
		if err != nil {
			// app.Logger.Error("merchant ID is required", err)
			utils.SendResponse(w, http.StatusBadRequest, "merchant ID is required", nil)
			return
		}

		if err := payload.Validate(); err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		resp, err := client.UpdateMosque(ctx, &pb.UpdateMosqueRequest{
			Id:         mosqueId.String(),
			Name:       payload.Name,
			Address:    payload.Address,
			EidTime:    &timestamppb.Timestamp{Seconds: payload.EidTime.Unix()},
			JummahTime: &timestamppb.Timestamp{Seconds: payload.JummahTime.Unix()},
			Lat:        payload.Lat,
			Lng:        payload.Lng,
		})

		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}
		// json.NewEncoder(w).Encode(resp)
		utils.SendResponse(w, http.StatusOK, "Mosque updated successfully", resp)
	}
}

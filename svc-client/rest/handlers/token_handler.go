package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jimeux/app-mesher/svc-client/rest/requests"
	"github.com/Jimeux/app-mesher/svc-client/rest/responses"
	"github.com/Jimeux/app-mesher/svc-client/rpc"
)

type TokenHandler struct {
	Handler
	identitySvc rpc.IdentityServiceClient
	profileSvc  rpc.ProfileServiceClient
}

func NewTokenHandler(identitySvc rpc.IdentityServiceClient, profileSvc rpc.ProfileServiceClient) *TokenHandler {
	return &TokenHandler{
		identitySvc: identitySvc,
		profileSvc:  profileSvc,
	}
}

func (h *TokenHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	var req requests.TokenGet
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, r, err)
		return
	}

	identityReply, err := h.identitySvc.IssueToken(r.Context(), &rpc.IssueTokenRequest{
		Username: req.Username,
	})
	if err != nil {
		h.writeError(w, r, err)
		return
	}

	profileReply, err := h.profileSvc.GetProfile(r.Context(), &rpc.GetProfileRequest{
		Id: req.ID,
	})
	if err != nil {
		h.writeError(w, r, err)
		return
	}

	h.writeJSON(w, r, responses.TokenGet{
		Token:   identityReply.GetToken(),
		Profile: profileReply.GetProfile(),
		Data:    profileReply.GetData(),
	})
}

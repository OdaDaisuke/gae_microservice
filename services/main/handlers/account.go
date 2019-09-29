package handler

import (
	"context"
	"fmt"
	accountpb "github.com/OdaDaisuke/gae_sand/pb/account"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountRoute struct {
	SigninHandler *SigninHandler

	accountClient accountpb.AccountServiceClient
	ctx context.Context
}

func NewAccountRoute(ac accountpb.AccountServiceClient, ctx context.Context) *AccountRoute {
	return &AccountRoute{
		SigninHandler: &SigninHandler{ac, ctx},
		accountClient: ac,
		ctx: ctx,
	}
}

/* ---------------
 * Children
 * --------------- */

/* /account/signin */

type SigninHandler struct {
	accountClient accountpb.AccountServiceClient
	ctx context.Context
}

func (sh SigninHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r.ParseForm()
	user, pw := r.Form["user"][0], r.Form["password"][0]

	sreq := &accountpb.SigninRequest{
		User: user,
		Password: pw,
	}
	sres, err := sh.accountClient.Signin(sh.ctx, sreq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "vars: %v\n", vars)
	fmt.Fprintf(w, "grpc response: %v\n", sres)

}

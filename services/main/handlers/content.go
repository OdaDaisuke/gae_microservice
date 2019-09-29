package handler

import (
	"context"
	"fmt"
	contentpb "github.com/OdaDaisuke/gae_sand/pb/content"
	"github.com/gorilla/mux"
	"net/http"
)

type ContentRoute struct {
	GetContentHandler *GetContentHandler
	GetContentsHandler *GetContentsHandler

	contentClient contentpb.ContentServiceClient
	ctx context.Context
}

func NewContentRoute(cc contentpb.ContentServiceClient, ctx context.Context) *ContentRoute {
	return &ContentRoute{
		GetContentHandler: &GetContentHandler{cc, ctx},
		GetContentsHandler: &GetContentsHandler{cc, ctx},
		contentClient: cc,
		ctx: ctx,
	}
}

/* ---------------
 * Children
 * --------------- */

/* /content */

type GetContentHandler struct {
	contentClient contentpb.ContentServiceClient
	ctx context.Context
}

func (ch GetContentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sreq := &contentpb.GetContentRequest{
		Id: id,
	}
	sres, err := ch.contentClient.GetContent(ch.ctx, sreq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "grpc response: %v\n", sres)
}

/* /contents */

type GetContentsHandler struct {
	contentClient contentpb.ContentServiceClient
	ctx context.Context
}

func (ch GetContentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	afterID, beforeID := vars["after_id"], vars["before_id"]

	creq := &contentpb.GetContentsRequest{}
	if afterID != "" {
		creq.AfterId = afterID
	} else {
		creq.BeforeId = beforeID
	}

	cres, err := ch.contentClient.GetContents(ch.ctx, creq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "grpc response: %v\n", cres)
}
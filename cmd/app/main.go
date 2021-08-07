package main

import (
	"bufio"
	"fmt"
	"github.com/dmitryrn/money/internal"
	"github.com/dmitryrn/money/internal/proto"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net/http"
)

func Article(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	url := queryValues.Get("q")
	if url == "" {
		http.Error(w, "Must specify the url to request", 400)
	}
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to retrieve article", 500)
	}
	if response.StatusCode >= 400 {
		if err != nil {
			http.Error(w, response.Status, response.StatusCode)
		}
	}
	reader := bufio.NewReader(response.Body)
	reader.WriteTo(w)
}

type GrpcWebMiddleware struct {
	*grpcweb.WrappedGrpcServer
}

func (m *GrpcWebMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsAcceptableGrpcCorsRequest(r) || m.IsGrpcWebRequest(r) {
			fmt.Println("somehting")
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewGrpcWebMiddleware(grpcWeb *grpcweb.WrappedGrpcServer) *GrpcWebMiddleware {
	return &GrpcWebMiddleware{grpcWeb}
}

func main() {
	grpcServer := grpc.NewServer()
	hackernewsService := internal.Server{}
	proto.RegisterBudgetingAppServer(grpcServer, hackernewsService)

	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		// Allow all origins, DO NOT do this in production
		return true
	}))

	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		NewGrpcWebMiddleware(wrappedGrpc).Handler,
	)

	router.Get("/article-proxy", Article)

	fmt.Println("Server is going to start on port 3001")

	if err := http.ListenAndServe(":3001", router); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}
}

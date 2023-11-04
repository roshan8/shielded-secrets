/*
Copyright © 2023 Roshan shetty roshan.aloor@gmail.com
*/
package cmd

import (
	"fmt"
	"net/http"

	"shielded-secrets/pkg/api"
	"shielded-secrets/vars"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "HTTP server for managing secrets",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		startApp()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// TODO: Rename the func name
func startApp() {
	vars.Init()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	startServer()
}

func startServer() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Use(middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
	)

	router.Get("/", api.ServeFrontendHandler)
	router.Route("/api", api.Routes)

	log.Info().Msg(fmt.Sprintf("starting webserver on port: %d ", vars.Port))
	log.Fatal().Msg(http.ListenAndServe(fmt.Sprintf(":%d", vars.Port), router).Error())
}

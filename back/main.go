package main

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/pansani/interjiajs-golang-example/database"
	"github.com/pansani/interjiajs-golang-example/models"
	"github.com/petaki/inertia-go"
)

// CustomFileServer sets the correct MIME type for JavaScript files
func CustomFileServer(root http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if filepath.Ext(r.URL.Path) == ".tsx" {
			w.Header().Set("Content-Type", "application/javascript")
		}
		http.FileServer(root).ServeHTTP(w, r)
	})
}

func homeHandler(inertiaManager *inertia.Inertia) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tasks []models.Task
		database.DB.Find(&tasks)

		data := map[string]interface{}{
			"name":  "Home Page",
			"tasks": tasks,
		}

		if err := inertiaManager.Render(w, r, "Home", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	url := "http://localhost:8080"
	rootTemplate := "../front/index.gohtml"
	version := ""

	inertiaManager := inertia.New(url, rootTemplate, version)

	mux := http.NewServeMux()
	mux.Handle("/", inertiaManager.Middleware(homeHandler(inertiaManager)))
	mux.Handle("/src/", http.StripPrefix("/src/", CustomFileServer(http.Dir("../front/src"))))

	// Set the MIME type for .js files explicitly
	mime.AddExtensionType(".js", "application/javascript")

	log.Println("Server is running at", url)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

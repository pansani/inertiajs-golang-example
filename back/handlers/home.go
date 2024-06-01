package handlers

import (
	"net/http"

	"github.com/petaki/inertia-go"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	inertiaManager := inertia.New("", "", "")

	data := map[string]interface{}{
		"tasks": []string{"Task 1", "Task 2", "Task 3"},
	}

	inertiaManager.Render(w, r, "Home", data)
}

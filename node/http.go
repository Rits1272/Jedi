package node

// TODO: locking mechanism
import (
  "fmt"
  "net/http"
  "encoding/json"
)

func getRouteHandler(w http.ResponseWriter, r *http.Request) {
  key := r.URL.Query().Get("key")
  exists, value := node.get(key)

  if !exists {
    http.Error(w, "Key not found!", http.StatusNotFound)
    return
  }

  response := map[string]string {"key": key, "value": value}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func setRouteHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
  }

  var request map[string]string

  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
      http.Error(w, "Invalid request payload", http.StatusBadRequest)
      return
  }

  key, keyExists := request["key"]
  value, valueExists := request["value"]

  if !keyExists || !valueExists {
    http.Error(w, "Missing key or value", http.StatusBadRequest)
    return
  }

  node.set(key, value)
  
  w.WriteHeader(http.StatusNoContent)
}

func mainRouteHandler(w http.ResponseWriter, r *http.Request) {
  response := map[string]string{"status": "jedi node is up and running!"}
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(response)
}

func StartServer() {
  http.HandleFunc("/", mainRouteHandler)
  http.HandleFunc("/get", getRouteHandler)
  http.HandleFunc("/set", setRouteHandler)

  port := "8080"
  fmt.Println("Staring server on port: ", port)

  if err := http.ListenAndServe(":"+port, nil); err != nil {
    fmt.Println("Unable to start server", err)
  }
}

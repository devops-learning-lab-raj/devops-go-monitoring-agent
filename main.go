package main
import ("encoding/json"; "log"; "net/http"; "os")
func main() {
  http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]string{"status":"ok","service":"monitoring-agent"})
  })
  port := os.Getenv("PORT"); if port == "" { port = "8080" }
  log.Printf("monitoring agent listening on :%s", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}

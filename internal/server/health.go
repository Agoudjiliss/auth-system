package server
import("net/http")

func ping(w http.ResponseWriter,r *http.Request){
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("pong\n"))
}

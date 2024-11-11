package red_api

import (
    "fmt"
    "net/http"
    "github.com/manjushettar/internal"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main(){
    router := chi.NewRouter()

    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)
    
    //subreddit endpoints
    router.Route("/subreddits", internal.subreddit)

    router.Route("/posts", internal.post)

    router.Route("/user", internal.user)
    
    fmt.Println("Server started.")
    http.ListenAndServe(":3333", router)
}

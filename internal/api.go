package internal

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)


func Subreddit(router chi.Router){
    router.Get("/", listSubreddits)
    router.Get("/{id}", getSubreddit)
    
    router.Get("/{id}/posts", getSubredditPosts)
    router.Post("/{id}/posts", addSubredditPosts)
    
    router.Post("/", createNewSubreddit)

    router.Post("/{id}/subscribe", addSubscription)
    router.Delete("/{id}/subscribe", deleteSubscription)
}


func listSubreddits(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    
    case http.MethodGet:
        subreddits, err := // get all subreddits from subreddit struct
        if err != nil{
            http.Error(w, "Something went wrong.", http.StatusInternalServerError)
            return
        }
        // write subreddits to client
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getSubreddit(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    
    case http.MethodGet:
        // get id 
        id, err := chi.URLParam(r, "id")

        if err != nil{
            http.Error(w, "Something went wrong.", http.StatusInternalServerError)
            return
        }

        subreddit, err := // fetch specific subreddit
        if err != nil{
            http.Error(w, "No subreddit found.", http.StatusNoContent)
            return
        }

        // write usbreddit to client
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func addSubredditPosts(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodPOST:
        /* request body:
        {
            "post":
        }
        */
        
        // decode subreddit id from req header

        // decode and validate req body

        // create new post

        // save to DB

        // write success if everything goes right, err at steps
    }
}

func createNewSubreddit(w http.ResponseWriter, r *http.Request){
    switch r.Method{

    case http.MethodPOST:
        /* Request body:
        {
            "name":
        }
        
        */ 
        // create a new subreddit from the req body
        
        // decode and validate the req body
        
        // create new subreddit

        // save to DB
        
        // write success if everything goes right, err between steps 
        
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func addSubscription(w http.ResponseWriter, r *http.Request){

    switch r.Method{

    case http.MethodPOST:
        // Request body:
        // { "user-id": "123" }
        
        // decode subreddit from req header

        // decode user_id from the req body
        
        // decode and validate the req body - look for userID
        
        // check if subreddit exists 

        // check if user is subscribed already

        // add subscription

        // save to DB
        
        // write success if everything goes right, err between steps 
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func deleteSubscription(w http.ResponseWriter, r *http.Request){
    switch r.Method {
    case http.MethodDELETE:
        // get subreddit id from req header

        // get user_id from auth context - can't get from req body

        // check if subreddit and user exist

        // check if user is subscribed

        // delete subscription from DB

        // write success if everything goes right, err between steps
    }
}

func getSubredditPosts(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    
    case http.MethodGet:
        // validate subreddit from req header

        // get all subreddit posts from subreddit struct
        
        //write posts to client using post struct
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}


func Post(router chi.Router){
    router.Get("/{id}", getPost)

    router.Post("/{id}/upvote", addUpvote)
    router.Delete("/{id}/upvote", removeUpvote)
    
    router.Get("/{id}/comments", getComments)
    router.Post("/{id}/comments", addComment)
}

func getPost(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodPOST:
        // decode and validate post from req header

        // verify post exists in DB

        // write post

        // return success
    }
}

func addUpvote(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodPOST:
        // decode and validate post from req header
        
        // verify post and userID exists

        // update user upvote history and post upvotes in DB

        // return success
    }
}

func removeUpvote(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodDELETE:
        // decode and validate post from req header

        // verify post and userID exists in upvote history

        // update user upvote history and post upvotes in DB

        // return success
    }
}

func getComments(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodGET:
        // decode and validate post id from req header

        // write all comments

        // return success
    }
}

func addComment(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodPOST:
        // decode and validate post id and userid from req header / auth

        // verify comment from req body

        // save comment to DB
        
        // return success
    }
}

func User(router chi.Router){
    router.Get("/", getProfile)
    router.Get("/{id}/subscribed", getSubscribed)
    router.Get("/{id}/upvoted", getUpvoted)
}

func getProfile(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodGET:
        // decode and validate userID from auth (JWT)
        
        // write information from user struct

        // return success
    }
}

func getSubscribed(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodGET:
        // decode and validate user ID from auth

        // get subscribed subreddits from DB

        // return success
    }
}

func getUpvoted(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodGET:
        // decode and validate user ID

        // get history of upvoted comments/posts from DB
        
        // return success
    }
}

package auth

import (
	"context"
	"log"
	globals "backend/globals"
	// helpers "backend/helpers"
)

func (r rpcInterface) AddWord(ctx context.Context, req *pb.AddWordRequest) (*pb.AddWordResponse, error) {
	f := r.raft.Apply([]byte(req.GetWord()), time.Second)
	if err := f.Error(); err != nil {
		return nil, rafterrors.MarkRetriable(err)
	}
	return &pb.AddWordResponse{
		CommitIndex: f.Index(),
	}, nil
}

func (r rpcInterface) GetWords(ctx context.Context, req *pb.GetWordsRequest) (*pb.GetWordsResponse, error) {
	r.wordTracker.mtx.RLock()
	defer r.wordTracker.mtx.RUnlock()
	return &pb.GetWordsResponse{
		BestWords:   cloneWords(r.wordTracker.words),
		ReadAtIndex: r.raft.AppliedIndex(),
	}, nil
}

func SignUp(username string, password string) string {
	globals.UserPass[username] = password
	return username
}

func SignIn(username string, password string) bool {
	return CheckUserPass(username, password)
}

func CheckUserPass(username, password string) bool {
	userpass := globals.UserPass

	log.Println("checkUserPass", username, password, userpass)

	if val, ok := userpass[username]; ok {
		log.Println(val, ok)
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
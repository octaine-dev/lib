package lib

import "fmt" // This will be referenced from the dynamically linked libstd

type Bet struct {
	amount int
	value string
}

type PlayRequest struct {
	Bet Bet
	State interface{}
	Payload interface{}
}

type PlayResponse struct {
	TotalWin int
	RoundEnd	 bool
	State interface{}
	Outcome interface{}
}

type RNG interface{
	Next(query []int) ([]int, error)
}


func Play(req *PlayRequest, rng RNG) (*PlayResponse, error) {
	rngResult, err := rng.Next([]int{1,2,3})
	if err != nil {
		return nil, err
	}
	fmt.Println("rng.Next(): ", rngResult)

	return &PlayResponse{
		100,
		true,
		[]string{"random", "engine", "data"},
		[]string{"outcome", "data"},
	}, nil

}



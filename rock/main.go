package rock

import (
	"net"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
<<<<<<< HEAD
<<<<<<< HEAD
	// "github.com/Saemkenya/rockscissorspaper_grpc/proto/rps.pb"
=======
	"github.com/Saemkenya/rockscissorspaper_grpc/proto/rps.pb"
>>>>>>> afd465589d7e2faea480ac200553d7e8f827d904
=======
	"github.com/Saemkenya/rockscissorspaper_grpc/proto/rps.pb"
>>>>>>> afd465589d7e2faea480ac200553d7e8f827d904
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":18701")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterRPSServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func randomChoice() string {
	rand.Seed(time.Now().UnixNano())
	options := [...]string{rock, paper, scissors}
	return options[rand.Intn(len(options))]
}

func (s *server) ChooseRock(ctx context.Context, req *Choice) (*OutCome, error) {
	computed := randomChoice()
	switch computed {
	case rock:
		result := pb.Medal_Charm
		winner := req.player
		looser := req.player
	case paper:
		result := pb.Medal_NoMedal
		winner := pb.Player{
			Name: "Computer Won!",
			Lives: 0,
			Medals: [pb.Medal_NoMedal],
			Wins: 0,
		}
		looser := req.player
	case paper:
		result := pb.Medal_Gold
		winner := req.player
		looser := pb.Player{
			Name: "Computer Lost!",
			Lives: 0,
			Medals: pb.[Medal_NoMedal],
			Wins: 0,
		}
	default:
		result := pb.Medal_NoMedal
		winner := pb.Player{
			Name: "No winner",
			Lives: 0,
			Medals: [pb.Medal_NoMedal]
			Wins: 0,		
		}
		looser := pb.Player{
			Name: "No looser",
			Lives: 0,
			Medals: [pb.Medal_NoMedal],
			Wins: 0,
		}
	}
	return &pb.OutCome{
		Result:	result,
		Winner: winner,
		Looser: looser,
	}
}
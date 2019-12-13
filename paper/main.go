package paper

import (
	"net"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/Saemkenya/rockscissorspaper_grpc/proto/rps.pb"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":18702")
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

func (s *server) ChoosePaper(ctx context.Context, req *Choice) (*OutCome, error) {
	computed := randomChoice()
	switch computed {
	case paper:
		result := pb.Medal_Charm
		winner := req.player
		looser := req.player
	case rock:
		result := pb.Medal_NoMedal
		winner := pb.Player{
			Name: "Computer Won!"
			Lives: 0,
			Medals: [pb.Medal_NoMedal],
			Wins: 0,
		}
		looser := req.player
	case scissors:
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
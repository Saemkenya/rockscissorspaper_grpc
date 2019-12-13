package main

import (
	"strconv"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"github.com/Saemkenya/rockscissorspaper_grpc/proto/rps.pb"
)

func main() {
	paperConn, err := grpc.Dial("localhost:18702", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	paperClient := pb.NewRPSServiceClient(paperConn)

	g := gin.Default()

	g.GET("/choose/:choice/:name", func(ctx *gin.Context){
		choice, err := ctx.Param("choice")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid choice parameter in argument."})
			return
		}
		name, nErr := ctx.Param("name")
		if nErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid name parameter in arg"})
		}
		switch choice {
		case rock:
			req := &pb.OutCome{
				Result: result,
				Winner: winner,
				Looser: looser,
			}
		}
		case paper:
			req := &pb.OutCome{
				Result: result,
				Winner: winner,
				Looser: looser,
			}
			choice := pb.Choice{
				Chosen: "paper",
				Player: pb.Player{
					Name: name,
					Lives: 10,
					Medals: 0,
					Wins: 0,
				}
			}
			if response, err := paperClient.ChoosePaper(ctx, choice); err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"result": fmt.Sprint(response.Result)
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":"Unable to call the paper service"
				})
			}
		case scissors:
			req := &pb.OutCome{
				Result: result,
				Winner: winner,
				Looser: looser,
			}
		default:
		}
	})

	// g.GET("/rps_arena", func(ctx *gin.Context){
	// })

	if err := g.Run(":18700"); err != nil {
		log.Fatalf("Failed to start server bcuz : [%v]", err)
	}

	req := &pb.OutCome{
		Medal
		Player
		Player
	}
}
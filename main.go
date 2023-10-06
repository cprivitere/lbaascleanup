package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/cprivitere/lbaascleanup/infrastructure"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("METAL_AUTH_TOKEN")
	if apiKey == "" {
		fmt.Println("METAL_AUTH_TOKEN environment variable not set")
		return
	}

	projectID := os.Getenv("METAL_PROJECT_ID")
	if apiKey == "" {
		fmt.Println("METAL_PROJECT_ID environment variable not set")
		return
	}

	lbClean := flag.Bool("lbClean", false, "delete all loadbalancers in the specified metro")
	poolClean := flag.Bool("poolClean", false, "delete all pools in the specified metro")
	metro := flag.String("metro", "da", "metro to clean up")
	flag.Parse()

	manager := infrastructure.NewManager(apiKey, projectID, *metro)

	if *lbClean {
		lbList, _ := manager.GetLoadBalancers(ctx)
		for _, j := range lbList.Loadbalancers {
			fmt.Println("Deleting: ", j.Name)
			_, err := manager.DeleteLoadBalancer(ctx, j.Id)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	if *poolClean {
		poolList, _ := manager.GetPools(ctx)
		for _, j := range poolList.Pools {
			fmt.Println("Deleting: ", j.Name)
			_, err := manager.DeleteLoadBalancerPool(ctx, j.Id)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

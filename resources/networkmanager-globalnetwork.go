package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/networkmanager"
)

type NetworkManagerGlobalNetwork struct {
	svc *networkmanager.NetworkManager
	ID  *string
}

func init() {
	register("NetworkManagerGlobalNetwork", ListNetworkManagerGlobalNetworks)
}

func ListNetworkManagerGlobalNetworks(sess *session.Session) ([]Resource, error) {
	svc := networkmanager.New(sess)
	resources := []Resource{}

	input := &networkmanager.DescribeGlobalNetworksInput{
		MaxResults: aws.Int64(100),
	}

	for {
		output, err := svc.DescribeGlobalNetworks(input)
		if err != nil {
			return nil, err
		}

		for _, globalNetwork := range output.GlobalNetworks {
			resources = append(resources, &NetworkManagerGlobalNetwork{
				svc: svc,
				ID:  globalNetwork.GlobalNetworkId,
			})
		}

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	return resources, nil
}

func (f *NetworkManagerGlobalNetwork) Remove() error {
	_, err := f.svc.DeleteGlobalNetwork(&networkmanager.DeleteGlobalNetworkInput{
		GlobalNetworkId: f.ID,
	})

	return err
}

func (f *NetworkManagerGlobalNetwork) String() string {
	return *f.ID
}

package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/networkmanager"
)

type NetworkManagerCoreNetwork struct {
	svc *networkmanager.NetworkManager
	ID  *string
}

func init() {
	register("NetworkManagerCoreNetwork", ListNetworkManagerCoreNetworks)
}

func ListNetworkManagerCoreNetworks(sess *session.Session) ([]Resource, error) {
	svc := networkmanager.New(sess)
	resources := []Resource{}

	input := &networkmanager.ListCoreNetworksInput{
		MaxResults: aws.Int64(100),
	}

	for {
		output, err := svc.ListCoreNetworks(input)
		if err != nil {
			return nil, err
		}

		for _, coreNetwork := range output.CoreNetworks {
			resources = append(resources, &NetworkManagerCoreNetwork{
				svc: svc,
				ID:  coreNetwork.CoreNetworkId,
			})
		}

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	return resources, nil
}

func (f *NetworkManagerCoreNetwork) Remove() error {
	_, err := f.svc.DeleteCoreNetwork(&networkmanager.DeleteCoreNetworkInput{
		CoreNetworkId: f.ID,
	})

	return err
}

func (f *NetworkManagerCoreNetwork) String() string {
	return *f.ID
}

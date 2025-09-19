package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/networkmanager"
)

type NetworkManagerConnectPeer struct {
	svc *networkmanager.NetworkManager
	ID  *string
}

func init() {
	register("NetworkManagerConnectPeer", ListNetworkManagerConnectPeers)
}

func ListNetworkManagerConnectPeers(sess *session.Session) ([]Resource, error) {
	svc := networkmanager.New(sess)
	resources := []Resource{}

	input := &networkmanager.ListConnectPeersInput{
		MaxResults: aws.Int64(100),
	}

	for {
		output, err := svc.ListConnectPeers(input)
		if err != nil {
			return nil, err
		}

		for _, connectPeer := range output.ConnectPeers {
			resources = append(resources, &NetworkManagerConnectPeer{
				svc: svc,
				ID:  connectPeer.ConnectPeerId,
			})
		}

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	return resources, nil
}

func (f *NetworkManagerConnectPeer) Remove() error {
	_, err := f.svc.DeleteConnectPeer(&networkmanager.DeleteConnectPeerInput{
		ConnectPeerId: f.ID,
	})

	return err
}

func (f *NetworkManagerConnectPeer) String() string {
	return *f.ID
}

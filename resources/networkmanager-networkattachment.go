package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/networkmanager"
)

type NetworkManagerNetworkAttachment struct {
	svc *networkmanager.NetworkManager
	ID  *string
}

func init() {
	register("NetworkManagerNetworkAttachment", ListNetworkManagerNetworkAttachments)
}

func ListNetworkManagerNetworkAttachments(sess *session.Session) ([]Resource, error) {
	svc := networkmanager.New(sess)
	resources := []Resource{}

	input := &networkmanager.ListAttachmentsInput{
		MaxResults: aws.Int64(100),
	}

	for {
		output, err := svc.ListAttachments(input)
		if err != nil {
			return nil, err
		}

		for _, NetworkAttachment := range output.Attachments {
			resources = append(resources, &NetworkManagerNetworkAttachment{
				svc: svc,
				ID:  NetworkAttachment.AttachmentId,
			})
		}

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	return resources, nil
}

func (f *NetworkManagerNetworkAttachment) Remove() error {
	_, err := f.svc.DeleteAttachment(&networkmanager.DeleteAttachmentInput{
		AttachmentId: f.ID,
	})

	return err
}

func (f *NetworkManagerNetworkAttachment) String() string {
	return *f.ID
}

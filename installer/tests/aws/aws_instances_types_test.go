package regression

import (
	"testing"
	testutil "tectonic/installer/tests/testutils"
	"github.com/aws/aws-sdk-go/aws"
)

func TestGetAwsInstancesTypes(t *testing.T) {

	// call to get the EC2 AWS Instance response
	resp, err := testutil.GetAwsInstances()

	if err != nil {
		t.Fatalf("there was an error listing instances in %s", err.Error())
	}

	for idx, res := range resp.Reservations {
		t.Logf(" Reservation Id: %s and Num Instances: %d ", *res.ReservationId, len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			t.Logf("    - Instance Type: %s", *inst.InstanceType)
			if len(*inst.InstanceType) == 0 {
				t.Fatalf("Instance type is empty for %s", *res.ReservationId)
			}
		}
	}
}

func TestGetAwsVolumes(t *testing.T) {

	//call to get the EC2 AWS Volumes

	resp, err := testutil.GetAwsVolumes()

	if err != nil {
		t.Fatalf("there was an error listing volumes in %s", err.Error())
	}

	for _, vol := range resp.Volumes {
		volumeId := aws.StringValue(vol.VolumeId)
		if len(volumeId) == 0 {
			t.Fatalf("volumeId is empty for %s", volumeId)
		}
		size := aws.Int64Value(vol.Size)
		if size == 0 {
			t.Fatalf("size is 0 for %s", volumeId)
		}
		iops := aws.Int64Value(vol.Iops)
		if iops == 0 {
			t.Fatalf("iops is 0 for %s", volumeId)
		}
		volumeType := aws.StringValue(vol.VolumeType)
		if len(volumeType) == 0 {
			t.Fatalf("volumeType is empty for %s", volumeId)
		}

		t.Logf("VolumeId: %s ", volumeId)
		t.Logf("size: %d ", size)
		t.Logf("iops: %d", iops)
		t.Logf("volumeType: %s", volumeType)

	}
}

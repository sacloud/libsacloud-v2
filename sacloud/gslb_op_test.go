package sacloud

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/naked"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

func TestGSLBOpCRUD(t *testing.T) {
	Test(t, &CRUDTestCase{
		SetupAPICaller: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testGSLBCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createGSLBExpected,
				IgnoreFields: ignoreGSLBFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testGSLBRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createGSLBExpected,
				IgnoreFields: ignoreGSLBFields,
			},
		},

		Update: &CRUDTestFunc{
			Func: testGSLBUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateGSLBExpected,
				IgnoreFields: ignoreGSLBFields,
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testGSLBDelete,
		},
	})
}

var (
	ignoreGSLBFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"FQDN",
		"Icon",
		"CreatedAt",
		"ModifiedAt",
	}
	createGSLBParam = &GSLBCreateRequest{
		Name:                    "libsacloud-v2-gslb",
		Description:             "desc",
		Tags:                    []string{"tag1", "tag2"},
		HealthCheckProtocol:     "http",
		HealthCheckHostHeader:   "usacloud.jp",
		HealthCheckPath:         "/index.html",
		HealthCheckResponseCode: types.StringNumber(200),
		DelayLoop:               20,
		Weighted:                types.StringTrue,
		SorryServer:             "8.8.8.8",
		DestinationServers: []*naked.GSLBServer{
			{
				IPAddress: "192.2.0.1",
				Enabled:   types.StringTrue,
			},
			{
				IPAddress: "192.2.0.2",
				Enabled:   types.StringTrue,
			},
		},
	}
	createGSLBExpected = &GSLB{
		Name:                    createGSLBParam.Name,
		Description:             createGSLBParam.Description,
		Tags:                    createGSLBParam.Tags,
		Availability:            types.Availabilities.Available,
		DelayLoop:               createGSLBParam.DelayLoop,
		Weighted:                createGSLBParam.Weighted,
		HealthCheckProtocol:     createGSLBParam.HealthCheckProtocol,
		HealthCheckHostHeader:   createGSLBParam.HealthCheckHostHeader,
		HealthCheckPath:         createGSLBParam.HealthCheckPath,
		HealthCheckResponseCode: createGSLBParam.HealthCheckResponseCode,
		HealthCheckPort:         createGSLBParam.HealthCheckPort,
		SorryServer:             createGSLBParam.SorryServer,
		DestinationServers: []*naked.GSLBServer{
			{
				IPAddress: "192.2.0.1",
				Enabled:   types.StringTrue,
			},
			{
				IPAddress: "192.2.0.2",
				Enabled:   types.StringTrue,
			},
		},
	}
	updateGSLBParam = &GSLBUpdateRequest{
		Name:                    "libsacloud-v2-gslb-upd",
		Description:             "desc-upd",
		Tags:                    []string{"tag1-upd", "tag2-upd"},
		HealthCheckProtocol:     "https",
		HealthCheckHostHeader:   "upd.usacloud.jp",
		HealthCheckPath:         "/index-upd.html",
		HealthCheckResponseCode: types.StringNumber(201),
		DelayLoop:               21,
		Weighted:                types.StringTrue,
		SorryServer:             "8.8.4.4",
		DestinationServers: []*naked.GSLBServer{
			{
				IPAddress: "192.2.1.1",
				Enabled:   types.StringTrue,
				Weight:    types.StringNumber(1000),
			},
			{
				IPAddress: "192.2.1.2",
				Enabled:   types.StringTrue,
				Weight:    types.StringNumber(1000),
			},
		},
	}
	updateGSLBExpected = &GSLB{
		Name:                    updateGSLBParam.Name,
		Description:             updateGSLBParam.Description,
		Tags:                    updateGSLBParam.Tags,
		Availability:            types.Availabilities.Available,
		DelayLoop:               updateGSLBParam.DelayLoop,
		Weighted:                updateGSLBParam.Weighted,
		HealthCheckProtocol:     updateGSLBParam.HealthCheckProtocol,
		HealthCheckHostHeader:   updateGSLBParam.HealthCheckHostHeader,
		HealthCheckPath:         updateGSLBParam.HealthCheckPath,
		HealthCheckResponseCode: updateGSLBParam.HealthCheckResponseCode,
		HealthCheckPort:         updateGSLBParam.HealthCheckPort,
		SorryServer:             updateGSLBParam.SorryServer,
		DestinationServers: []*naked.GSLBServer{
			{
				IPAddress: "192.2.1.1",
				Enabled:   types.StringTrue,
				Weight:    types.StringNumber(1000),
			},
			{
				IPAddress: "192.2.1.2",
				Enabled:   types.StringTrue,
				Weight:    types.StringNumber(1000),
			},
		},
	}
)

func testGSLBCreate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewGSLBOp(caller)
	return client.Create(context.Background(), DefaultZone, createGSLBParam)
}

func testGSLBRead(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewGSLBOp(caller)
	return client.Read(context.Background(), DefaultZone, testContext.ID)
}

func testGSLBUpdate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewGSLBOp(caller)
	return client.Update(context.Background(), DefaultZone, testContext.ID, updateGSLBParam)
}

func testGSLBDelete(testContext *CRUDTestContext, caller APICaller) error {
	client := NewGSLBOp(caller)
	return client.Delete(context.Background(), DefaultZone, testContext.ID)
}

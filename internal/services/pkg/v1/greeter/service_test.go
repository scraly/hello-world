package greeter_test

// import (
// 	"context"
// 	"testing"

// 	. "github.com/onsi/gomega"
// 	"go.uber.org/zap"
// 	"go.zenithar.org/pkg/log"

// 	"code.int.be.continental.cloud/ehorizon/services.nds.schema/internal/services/pkg/v1/db"
// 	dbv1 "code.int.be.continental.cloud/ehorizon/services.nds.schema/pkg/protocol/schema/db/v1"
// )

// func TestSayHello(t *testing.T) {
// 	ctx := context.Background()

// 	testCases := []struct {
// 		name        string
// 		req         *dbv1.GetInitDBRequest
// 		basePath    string
// 		fileName    string
// 		expectedErr bool
// 		checksum    string
// 		length      int64
// 	}{
// 		{
// 			name:        "req nil",
// 			req:         nil,
// 			basePath:    "",
// 			fileName:    "",
// 			expectedErr: true,
// 		},
// 		{
// 			name:        "input file not exists",
// 			req:         &dbv1.GetInitDBRequest{},
// 			basePath:    "/tmp/",
// 			fileName:    "toto.tgz",
// 			expectedErr: true,
// 		},
// 		{
// 			name:        "output file good",
// 			req:         &dbv1.GetInitDBRequest{},
// 			basePath:    "../../../../../tests/",
// 			fileName:    "NDS_INIT_DB.tgz",
// 			expectedErr: false,
// 			checksum:    "cf54b35476588f146c63c090415a125d3da91dc4db2faafcc20b75ee88f0e192cf0ed8d71282a4d443718e4f4c5284c0149c82887ddd1c7b1829c01522e48aaf",
// 			length:      1416509,
// 		},
// 		{
// 			name:        "input file empty",
// 			req:         &dbv1.GetInitDBRequest{},
// 			basePath:    "../../../../../tests/",
// 			fileName:    "ERROR.tgz",
// 			expectedErr: true,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		testCase := tc
// 		t.Run(testCase.name, func(t *testing.T) {
// 			t.Parallel()
// 			g := NewGomegaWithT(t)

// 			underTest := db.New(testCase.basePath, testCase.fileName)

// 			got, err := underTest.GetInitDB(ctx, testCase.req)

// 			// assert errors and results expectations
// 			if testCase.expectedErr {
// 				log.For(ctx).Info("error expected", zap.Any("err", got.Error))
// 				g.Expect(got).ToNot(BeNil(), "Result should be nil")
// 				g.Expect(got.Error.Code).ToNot(BeNil(), "Error code should not be nil")
// 				// g.Expect(got.Error).ToNot(BeNil(), "Error should not be nil")
// 				g.Expect(err).To(BeNil(), "Error not be nil")
// 			} else {
// 				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
// 				g.Expect(got.Error).To(BeNil(), "Errors should be nil")

// 				// Test checksum and length
// 				g.Expect(got.Entity.CheckSum).To(Equal(testCase.checksum))
// 				g.Expect(got.Entity.Length).To(Equal(testCase.length))
// 			}
// 		})
// 	}
// }

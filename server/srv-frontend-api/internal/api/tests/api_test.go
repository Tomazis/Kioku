package api_test

import (
	"context"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tomazis/kioku/server/srv-frontend-api/internal/api"
	"github.com/tomazis/kioku/server/srv-frontend-api/internal/logger"
	mock_api "github.com/tomazis/kioku/server/srv-frontend-api/internal/mocks"
	pb "github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var ctrl *gomock.Controller
var r *mock_api.MockRepo

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

var _ = BeforeEach(func() {
	// var err error
	// Expect(err).ShouldNot(HaveOccurred())
	ctrl = gomock.NewController(GinkgoT())
	r = mock_api.NewMockRepo(ctrl)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	// r := repo.NewRepo()
	pb.RegisterSrvFrontendApiServiceServer(s, api.NewFrontendAPI(r))
	go func() {
		if err := s.Serve(lis); err != nil {
			logger.FatalKV(context.Background(), "Server exited with error", "error", err)
		}
	}()
})

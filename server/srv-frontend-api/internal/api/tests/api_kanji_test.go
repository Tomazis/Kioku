package api_test

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	models "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/kanji"
	pb "github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api"
	"google.golang.org/grpc"
)

var kanjiList = []rune("万丈三上下丸久亡凡刃千口土士夕大女子寸小山川工己干弓才之巾乞于也々勺")

var _ = Describe("test Kanji API get method", func() {
	rand.Seed(time.Now().UnixNano())

	kanjiAmount := 20

	var kanjiContainer []*models.Kanji
	ctx := context.Background()

	for i := 1; i < kanjiAmount; i++ {
		kanjiContainer = append(kanjiContainer, &models.Kanji{
			ID:           uint64(i),
			Kanji:        string(kanjiList[rand.Intn(len(kanjiList))]),
			Primary:      "a",
			Level:        uint32(rand.Intn(kanjiAmount)),
			Alternatives: nil,
			Onyomi:       nil,
			Kunyomi:      nil,
		})
	}

	retrieve_kanji := func(id uint64) (*models.Kanji, error) {
		for _, kan := range kanjiContainer {
			if kan.ID == id {
				return kan, nil
			}
		}
		return nil, errors.New("does not have that id")
	}
	When("DB table has kanji models", func() {
		Context("db table has that kanji id", func() {
			It("should return kanji model by ID", func() {
				kanji_id := uint64(rand.Intn(len(kanjiContainer))) + 1
				r.EXPECT().GetKanji(gomock.Any(), gomock.Any()).Return(retrieve_kanji(kanji_id)).AnyTimes()

				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
				Expect(err).ShouldNot(HaveOccurred())
				defer conn.Close()
				client := pb.NewSrvFrontendApiServiceClient(conn)
				resp, err := client.GetKanjiV1(ctx, &pb.GetKanjiV1Request{KanjiId: kanji_id})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.Kanji.Id).To(Equal(kanji_id))

			})
		})
		Context("db table does not have that kanji id", func() {
			It("should return nil and error", func() {
				kanji_id := uint64(len(kanjiContainer) + 10)
				r.EXPECT().GetKanji(gomock.Any(), gomock.Any()).Return(retrieve_kanji(kanji_id)).AnyTimes()

				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
				Expect(err).ShouldNot(HaveOccurred())
				defer conn.Close()
				client := pb.NewSrvFrontendApiServiceClient(conn)
				resp, err := client.GetKanjiV1(ctx, &pb.GetKanjiV1Request{KanjiId: kanji_id})

				Expect(err).Should(HaveOccurred())
				Expect(resp).Should(BeNil())
			})
		})

	})
	When("problem with db service", func() {
		It("should return nil and error", func() {
			kanji_id := uint64(len(kanjiContainer) + 10)
			r.EXPECT().GetKanji(gomock.Any(), gomock.Any()).Return(nil, errors.New("can not establish connection")).AnyTimes()

			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			Expect(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			client := pb.NewSrvFrontendApiServiceClient(conn)
			resp, err := client.GetKanjiV1(ctx, &pb.GetKanjiV1Request{KanjiId: kanji_id})

			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})
	When("problem with validation", func() {
		It("should return nil and error", func() {
			kanji_id := uint64(0)

			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			Expect(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			client := pb.NewSrvFrontendApiServiceClient(conn)

			resp, err := client.GetKanjiV1(ctx, &pb.GetKanjiV1Request{KanjiId: kanji_id})
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())

		})
	})
})

var _ = Describe("test Kanji API list method", func() {
	rand.Seed(time.Now().UnixNano())

	kanjiAmount := 20

	var kanjiContainer []*models.Kanji
	ctx := context.Background()

	for i := 1; i < kanjiAmount*2; i += 2 {
		kanjiContainer = append(kanjiContainer, &models.Kanji{
			ID:           uint64(i),
			Kanji:        string(kanjiList[rand.Intn(len(kanjiList))]),
			Primary:      "a",
			Level:        uint32(i),
			Alternatives: nil,
			Onyomi:       nil,
			Kunyomi:      nil,
		})
		kanjiContainer = append(kanjiContainer, &models.Kanji{
			ID:           uint64(i + 1),
			Kanji:        string(kanjiList[rand.Intn(len(kanjiList))]),
			Primary:      "a",
			Level:        uint32(i),
			Alternatives: nil,
			Onyomi:       nil,
			Kunyomi:      nil,
		})
	}

	retrieve_list := func(level uint32) ([]*models.Kanji, error) {
		var ret []*models.Kanji
		for _, kan := range kanjiContainer {
			if kan.Level == level {
				ret = append(ret, kan)
			}
		}
		if len(ret) == 0 {
			return nil, errors.New("Empty list")
		}
		return ret, nil
	}

	When("DB table has kanji models", func() {
		Context("db table has that kanji with level", func() {
			It("should return kanji list by level", func() {
				kanji_level := uint32(rand.Intn(kanjiAmount)) + 1
				r.EXPECT().ListKanji(gomock.Any(), gomock.Any()).Return(retrieve_list(kanji_level)).AnyTimes()

				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
				Expect(err).ShouldNot(HaveOccurred())
				defer conn.Close()
				client := pb.NewSrvFrontendApiServiceClient(conn)
				resp, err := client.ListKanjiV1(ctx, &pb.ListKanjiV1Request{Level: kanji_level})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.Kanji).ShouldNot(BeEmpty())
				for _, k := range resp.Kanji {
					Expect(k.Level).To(Equal(kanji_level))
				}
			})
		})
		Context("db table does not have that kanji level", func() {
			It("should return nil and error", func() {
				kanji_level := uint32(len(kanjiContainer) + 10)
				r.EXPECT().ListKanji(gomock.Any(), gomock.Any()).Return(retrieve_list(kanji_level)).AnyTimes()

				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
				Expect(err).ShouldNot(HaveOccurred())
				defer conn.Close()
				client := pb.NewSrvFrontendApiServiceClient(conn)
				resp, err := client.ListKanjiV1(ctx, &pb.ListKanjiV1Request{Level: kanji_level})

				Expect(err).Should(HaveOccurred())
				Expect(resp).Should(BeNil())
			})
		})

	})
	When("problem with db service", func() {
		It("should return nil and error", func() {
			kanji_level := uint32(len(kanjiContainer))
			r.EXPECT().ListKanji(gomock.Any(), gomock.Any()).Return(nil, errors.New("can not establish connection")).AnyTimes()

			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			Expect(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			client := pb.NewSrvFrontendApiServiceClient(conn)
			resp, err := client.ListKanjiV1(ctx, &pb.ListKanjiV1Request{Level: kanji_level})

			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("problem with validation", func() {
		It("should return nil and error", func() {
			kanji_level := uint32(0)

			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			Expect(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			client := pb.NewSrvFrontendApiServiceClient(conn)
			resp, err := client.ListKanjiV1(ctx, &pb.ListKanjiV1Request{Level: kanji_level})

			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())

		})
	})
})

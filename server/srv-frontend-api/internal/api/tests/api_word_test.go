package api_test

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	m_word "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/word"
	pb "github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api"
	"google.golang.org/grpc"
)

var _ = Describe("test Word API get method", func() {
	rand.Seed(time.Now().UnixNano())

	wordAmount := 20

	var wordContainer []*m_word.Word

	ctx := context.Background()

	for i := 1; i < wordAmount; i++ {
		wordContainer = append(wordContainer, &m_word.Word{
			ID:           uint64(i),
			Word:         "hello",
			Primary:      "a",
			Level:        uint32(rand.Intn(wordAmount)),
			Composition:  nil,
			Alternatives: nil,
			Readings:     nil,
			Types:        nil,
			Sentences:    nil,
		})
	}

	retrieve_word := func(id uint64) (*m_word.Word, error) {
		for _, w := range wordContainer {
			if w.ID == id {
				return w, nil
			}
		}
		return nil, errors.New("does not have that id")
	}
	When("DB table has word models", func() {
		Context("db table has that word id", func() {
			It("should return word model by ID", func() {
				word_id := uint64(rand.Intn(len(wordContainer))) + 1
				r.EXPECT().GetWord(gomock.Any(), gomock.Any()).Return(retrieve_word(word_id)).AnyTimes()

				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
				Expect(err).ShouldNot(HaveOccurred())
				defer conn.Close()
				client := pb.NewSrvFrontendApiServiceClient(conn)
				resp, err := client.GetWordV1(ctx, &pb.GetWordV1Request{WordId: word_id})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.Word.Id).To(Equal(word_id))
			})
		})
		Context("db table does not have that word id", func() {
			It("should return nil and error", func() {
				word_id := uint64(len(wordContainer) + 10)
				r.EXPECT().GetWord(gomock.Any(), gomock.Any()).Return(retrieve_word(word_id)).AnyTimes()

				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
				Expect(err).ShouldNot(HaveOccurred())
				defer conn.Close()
				client := pb.NewSrvFrontendApiServiceClient(conn)
				resp, err := client.GetWordV1(ctx, &pb.GetWordV1Request{WordId: word_id})

				Expect(err).Should(HaveOccurred())
				Expect(resp).Should(BeNil())
			})
		})
	})
	When("problem with db service", func() {
		It("should return nil and error", func() {
			word_id := uint64(len(wordContainer) + 10)
			r.EXPECT().GetWord(gomock.Any(), gomock.Any()).Return(nil, errors.New("can not establish connection")).AnyTimes()

			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			Expect(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			client := pb.NewSrvFrontendApiServiceClient(conn)
			resp, err := client.GetWordV1(ctx, &pb.GetWordV1Request{WordId: word_id})

			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})
	When("problem with validation", func() {
		It("should return nil and error", func() {
			word_id := uint64(0)

			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			Expect(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			client := pb.NewSrvFrontendApiServiceClient(conn)

			resp, err := client.GetWordV1(ctx, &pb.GetWordV1Request{WordId: word_id})
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())

		})
	})

})

var _ = Describe("test Word API list method", func() {
	rand.Seed(time.Now().UnixNano())

	wordAmount := 20

	var wordContainer []*m_word.Word
	ctx := context.Background()

	for i := 1; i < wordAmount*2; i += 2 {
		wordContainer = append(wordContainer, &m_word.Word{
			ID:           uint64(i),
			Word:         "hello",
			Primary:      "a",
			Level:        uint32(rand.Intn(wordAmount)),
			Composition:  nil,
			Alternatives: nil,
			Readings:     nil,
			Types:        nil,
			Sentences:    nil,
		})
		wordContainer = append(wordContainer, &m_word.Word{
			ID:           uint64(i + 1),
			Word:         "hello",
			Primary:      "a",
			Level:        uint32(rand.Intn(wordAmount)),
			Composition:  nil,
			Alternatives: nil,
			Readings:     nil,
			Types:        nil,
			Sentences:    nil,
		})
	}

	retrieve_list := func(level uint32) ([]*m_word.Word, error) {
		var ret []*m_word.Word
		for _, kan := range wordContainer {
			if kan.Level == level {
				ret = append(ret, kan)
			}
		}
		if len(ret) == 0 {
			return nil, errors.New("Empty list")
		}
		return ret, nil
	}

	When("DB table has word models", func() {
		Context("db table has that word with level", func() {
			It("should return word list by level", func() {
				word_level := uint32(rand.Intn(wordAmount)) + 1
				r.EXPECT().ListWordsByLevel(gomock.Any(), gomock.Any()).Return(retrieve_list(word_level)).AnyTimes()

				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
				Expect(err).ShouldNot(HaveOccurred())
				defer conn.Close()
				client := pb.NewSrvFrontendApiServiceClient(conn)
				resp, err := client.ListWordsByLevelV1(ctx, &pb.ListWordsByLevelV1Request{Level: word_level})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.Words).ShouldNot(BeEmpty())
				for _, w := range resp.Words {
					Expect(w.Level).To(Equal(word_level))
				}
			})
		})
		Context("db table does not have that word level", func() {
			It("should return nil and error", func() {
				word_level := uint32(len(wordContainer) + 10)
				r.EXPECT().ListWordsByLevel(gomock.Any(), gomock.Any()).Return(retrieve_list(word_level)).AnyTimes()

				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
				Expect(err).ShouldNot(HaveOccurred())
				defer conn.Close()
				client := pb.NewSrvFrontendApiServiceClient(conn)
				resp, err := client.ListWordsByLevelV1(ctx, &pb.ListWordsByLevelV1Request{Level: word_level})

				Expect(err).Should(HaveOccurred())
				Expect(resp).Should(BeNil())
			})
		})
		// Context("db table has that word with kanji", func() {
		// 	It("should return word list by level", func() {
		// 		word_kanji := uint32(rand.Intn(wordAmount)) + 1
		// 		r.EXPECT().ListWordByKanji(gomock.Any(), gomock.Any()).Return(retrieve_list(word_level)).AnyTimes()

		// 		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
		// 		Expect(err).ShouldNot(HaveOccurred())
		// 		defer conn.Close()
		// 		client := pb.NewSrvFrontendApiServiceClient(conn)
		// 		resp, err := client.ListWordsByLevelV1(ctx, &pb.ListWordsByLevelV1Request{Level: word_level})

		// 		Expect(err).ShouldNot(HaveOccurred())
		// 		Expect(resp).ShouldNot(BeNil())
		// 		Expect(resp.Words).ShouldNot(BeEmpty())
		// 		for _, w := range resp.Words {
		// 			Expect(w.Level).To(Equal(word_level))
		// 		}
		// 	})
		// ))

	})
	When("problem with db service", func() {
		It("should return nil and error", func() {
			word_level := uint32(len(wordContainer))
			r.EXPECT().ListWordsByLevel(gomock.Any(), gomock.Any()).Return(nil, errors.New("can not establish connection")).AnyTimes()

			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			Expect(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			client := pb.NewSrvFrontendApiServiceClient(conn)
			resp, err := client.ListWordsByLevelV1(ctx, &pb.ListWordsByLevelV1Request{Level: word_level})

			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("problem with validation", func() {
		It("should return nil and error", func() {
			word_level := uint32(0)

			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			Expect(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			client := pb.NewSrvFrontendApiServiceClient(conn)
			resp, err := client.ListWordsByLevelV1(ctx, &pb.ListWordsByLevelV1Request{Level: word_level})

			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())

		})
	})
})

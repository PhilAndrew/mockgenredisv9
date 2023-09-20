package mock_v9

import (
	"context"
	"encoding/binary"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
	gomock "go.uber.org/mock/gomock"
)

var _ = Describe("Redis Cache V2", func() {
	var (
		mockCtrl  *gomock.Controller
		mockRedis *MockCmdable

		cache *RedisV2
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRedis = NewMockCmdable(mockCtrl)

		cache = NewRedisCacheV2(mockRedis).(*RedisV2)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Store Actions", func() {

		It("should successfully call MGet", func() {

			mockRedis.EXPECT().MGet(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(context context.Context, cacheKey string, deleteKey string) *redis.SliceCmd {
				defer GinkgoRecover()
				buf := make([]byte, 8)
				binary.BigEndian.PutUint64(buf, uint64(100))
				slice := append(make([]interface{}, 0), string("test"), buf)
				return redis.NewSliceResult(slice, nil)
			}).Times(1)

			err := cache.TestMGet(context.Background())
			Expect(err).To(BeNil())
		})
	})
})

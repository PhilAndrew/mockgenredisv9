
# Potential bug between mockgen and redis cache v9

## Introduction

Running `go test -cover ./...` in the root directory here or running the test suite in `mock_v9` directory shows the test failure result

`wrong number of arguments in DoAndReturn func for *mock_v9.MockCmdable.MGet: got 3, want 2 [/build/mock_v9/redis_cache_v2_test.go:36]`

This is not expected as 3 arguments are supplied and available on the method, I would expect this test which is in file `redis_cache_v2_test.go` to pass.

`mockRedis.EXPECT().MGet(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(context context.Context, cacheKey string, deleteKey string)`

## Mock file generation

To generate the mock files (which already exist).

`mockgen github.com/redis/go-redis/v9 Cmdable > mock_v9/redis_mock.go`

`mockgen github.com/redis/go-redis/v9 Pipeliner > mock_v9/redis_pipeliner_mock.go`


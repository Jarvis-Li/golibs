package redis

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

type RankItem struct {
	ID    string
	Score float64
}

func ZRange(Redis *pool.Pool, key string, start, end interface{}) (items []*RankItem, err error) {
	return zsetRange(Redis, key, start, end, false)
}

func ZRevRange(Redis *pool.Pool, key string, start, end interface{}) (items []*RankItem, err error) {
	return zsetRange(Redis, key, start, end, true)
}

func zsetRange(Redis *pool.Pool, key string, start, end interface{}, desc bool) (items []*RankItem, err error) {
	cmd := ""
	if desc {
		cmd = "ZREVRANGE"
	} else {
		cmd = "ZRANGE"
	}

	res, err := Redis.Cmd(cmd, key, start, end, "WITHSCORES").Array()
	if err != nil {
		return
	}

	l := len(res)
	for i := 0; i < l; i += 2 {
		id, err := res[i].Str()
		if err != nil {
			return nil, err
		}

		score, err := res[i+1].Float64()
		if err != nil {
			return nil, err
		}

		items = append(items, &RankItem{
			ID:    id,
			Score: score,
		})
	}

	return
}

func Rank(Redis *pool.Pool, key string, me string, n int) (items []*RankItem, myRank int, err error) {
	items, err = ZRevRange(Redis, key, 0, n-1)
	if err != nil {
		return
	}

	var hasMe bool
	for i, item := range items {
		if item.ID == me {
			hasMe = true
			myRank = i
		}
	}
	if hasMe {
		return
	}

	res := Redis.Cmd("ZREVRANK", key, me)
	if res.IsType(redis.Nil) {
		return
	}
	myRank, err = res.Int()
	if err != nil {
		return
	}
	myScore, err := Redis.Cmd("ZSCORE", key, me).Float64()
	if err != nil {
		return
	}
	items = append(items, &RankItem{
		ID:    me,
		Score: myScore,
	})
	return
}

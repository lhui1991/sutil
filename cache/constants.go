package cache

const (
	SpanLogKeyKey    = "key"
	SpanLogCacheType = "cache"
	SpanLogOp 		 = "op"
)

type CacheType int

const (
	CacheTypeRedis = iota
)

func (t CacheType) String() string {
	switch t {
	case CacheTypeRedis:
		return "redis"
	default:
		return ""
	}
}
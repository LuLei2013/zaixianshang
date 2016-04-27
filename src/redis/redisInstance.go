package redis



func GetRedisInstance() *RedisClient {
	redissvr := "localhost:6379"

	conntimeout := 100
	readtimeout := 50
	writetimeout := 50
	maxidle := 100
	maxactive := 1000
	expiresecond := 7000
	return newRedisClient(redissvr, conntimeout, readtimeout, writetimeout, maxidle, maxactive, expiresecond)
}

func newRedisClient(redissvr string, conntimeout, readtimeout, writetimeout, maxidle, maxactive, expiresecond int) *RedisClient {

	rc := new(RedisClient)
	if rc == nil {
		return nil
	}

	rc.pool = GetRedisPool(redissvr, conntimeout, readtimeout, writetimeout, maxidle, maxactive)
	if rc.pool == nil {
		return nil
	}

	rc.redissvr = redissvr

	rc.conntimeout = conntimeout
	rc.readtimeout = readtimeout
	rc.writetimeout = writetimeout
	rc.maxidle = maxidle
	rc.maxactive = maxactive
	rc.expiresecond = expiresecond

	return rc
}
package main

import (
	"fmt"
	"time"
)

// RateLimiter 생성: 특정 시간 동안 허용된 요청 횟수를 제한
func rateLimiter(limit int, interval time.Duration) func() bool {
	count := 0
	lastReset := time.Now() // 1초 경과 계산을 위해 타임스탬프

	return func() bool {
		now := time.Now()

		// 시간이 경과했으면 카운트 초기화
		if now.Sub(lastReset) > interval { // 위에서 찍은 타임스탬프와 현재시간 간격계산 > 1초(interval)
			lastReset = now // 타임스탬프를 리셋
			count = 0 // 카운트를 리셋
		}

		// 요청이 제한 내에 있으면 허용
		if count < limit { // 0,1,2,limit
			count++
			return true
		}

		// 요청 초과 시 거부
		return false
	}
}

func main() {
	// 초당 3개의 요청만 허용하는 RateLimiter 생성
	limiter := rateLimiter(3, time.Second)

	for i := 1; i <= 10; i++ {
		time.Sleep(200 * time.Millisecond) // 요청 간격 200ms
		if limiter() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Denied (Rate limit exceeded)\n", i)
		}
	}
}

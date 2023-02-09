package main

// https://leetcode.cn/problems/design-authentication-manager/
// 1797. 设计一个验证系统

// 需要一个验证码的有效时间字段
// 需要一个过期时间 map --> key 为 tokenId, value 为 过期的时间
type AuthenticationManager struct {
	defaultLiveTime     int
	tokenMapExpiredTime map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		defaultLiveTime:     timeToLive,
		tokenMapExpiredTime: make(map[string]int),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokenMapExpiredTime[tokenId] = currentTime + this.defaultLiveTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if expiredTime, ok := this.tokenMapExpiredTime[tokenId]; ok {
		if currentTime < expiredTime {
			this.Generate(tokenId, currentTime)
		} else {
			delete(this.tokenMapExpiredTime, tokenId)
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var res int
	expiredKeys := make([]string, 0)
	for key, val := range this.tokenMapExpiredTime {
		if val > currentTime {
			res++
		} else {
			expiredKeys = append(expiredKeys, key)
		}
	}
	for _, key := range expiredKeys {
		delete(this.tokenMapExpiredTime, key)
	}
	return res
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */

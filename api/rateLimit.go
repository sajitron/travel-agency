package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	defaultExpiryTime = 3600 // in seconds i.e. each token has a lifetime of 1 hour
	maximumTries      = 5    // bucket limit
)

var ctx = context.Background()

type FailedAccess struct {
	Identity   string
	AccessTime time.Time
	ExpiryTime time.Time
}

// rateLimit checks if the identity has failed to login more than a specific number of times in the last hour
func rateLimit(identity string, redisClient *redis.Client) error {
	loginIdentity := fmt.Sprintf("%s-login", identity)
	// check if store exists
	existingIdentity, err := redisClient.LRange(ctx, loginIdentity, 0, -1).Result()
	if err != nil {
		fmt.Println("Error getting list:", err)
		return nil
	}

	// check length
	listLength, err := redisClient.LLen(ctx, loginIdentity).Result()
	if listLength >= maximumTries {
		return fmt.Errorf("maximum tries attained, try again later")
	}

	if err != nil {
		fmt.Println("Error checking the length of list:", err)
		return nil
	}

	// loop over existing list and remove expired failed access
	for i := 0; i < int(listLength); i++ {
		failedAccess, _ := redisClient.LIndex(ctx, loginIdentity, int64(i)).Result()

		var structInList FailedAccess
		err = json.Unmarshal([]byte(failedAccess), &structInList)
		if err != nil {
			fmt.Println("Error unmarshalling failed access:", err)
			continue
		}

		if structInList.ExpiryTime.Before(time.Now()) {
			_, err := redisClient.LRem(ctx, loginIdentity, 0, failedAccess).Result()
			if err != nil {
				fmt.Println("Error removing expired failed access", err)
			} else {
				fmt.Printf("Expired failed access - %s - removed successfully", loginIdentity)
			}
		}
	}

	if existingIdentity == nil || listLength < maximumTries {
		expiryTime := defaultExpiryTime * time.Second
		currentTime := time.Now()
		failedAccess := FailedAccess{
			Identity:   identity,
			AccessTime: currentTime,
			ExpiryTime: currentTime.Add(time.Second * defaultExpiryTime),
		}

		jsonBytes, err := json.Marshal(failedAccess)
		if err != nil {
			fmt.Println("Error marshalling struct to JSON:", err)
			return nil
		}

		err = redisClient.RPush(ctx, loginIdentity, string(jsonBytes)).Err()
		if err != nil {
			fmt.Println("Error adding slice to new redis list:", err)
			return nil
		}

		err = redisClient.Expire(ctx, loginIdentity, expiryTime).Err()
		if err != nil {
			fmt.Println("Error setting expiry time on redis list:", err)
			return nil
		}
	}

	fmt.Printf("Failed access logged for %s", identity)
	return nil
}

// TODO: add tests

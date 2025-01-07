// my soluion, although chatgp pointed out that i need to use a big queue list not circular and cleaned up my code
package main

type Tweet struct {
	userID  int
	tweetID int
}

type userInfo struct {
	followList map[int]struct{}
}

type Twitter struct {
	tweets    []Tweet           // Global list of all tweets
	userInfos map[int]*userInfo // UserID to UserInfo
}

func Constructor() Twitter {
	return Twitter{
		tweets:    []Tweet{},
		userInfos: make(map[int]*userInfo),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	// Add tweet to global tweet list
	this.tweets = append(this.tweets, Tweet{userID: userId, tweetID: tweetId})

	// Ensure user exists in userInfos
	if _, ok := this.userInfos[userId]; !ok {
		this.userInfos[userId] = &userInfo{
			followList: make(map[int]struct{}),
		}
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	result := []int{}
	count := 0

	// Ensure user exists
	if _, ok := this.userInfos[userId]; !ok {
		this.userInfos[userId] = &userInfo{followList: make(map[int]struct{})}
	}

	// Traverse tweets in reverse order (latest first)
	for i := len(this.tweets) - 1; i >= 0 && count < 10; i-- {
		tweet := this.tweets[i]

		// Check if the tweet belongs to the user or one of their followees
		if tweet.userID == userId || this.isFollowing(userId, tweet.userID) {
			result = append(result, tweet.tweetID)
			count++
		}
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	// Ensure follower exists
	if _, ok := this.userInfos[followerId]; !ok {
		this.userInfos[followerId] = &userInfo{
			followList: make(map[int]struct{}),
		}
	}

	// Add followee to the follower's followList
	this.userInfos[followerId].followList[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if user, ok := this.userInfos[followerId]; ok {
		delete(user.followList, followeeId)
	}
}

func (this *Twitter) isFollowing(followerId, followeeId int) bool {
	if user, ok := this.userInfos[followerId]; ok {
		_, follows := user.followList[followeeId]
		return follows
	}
	return false
}

func test2() {
	twitter := Constructor()
	twitter.PostTweet(1, 5)   // User 1 posts a new tweet with id = 10.
	twitter.PostTweet(1, 3)   // User 2 posts a new tweet with id = 20.
	twitter.PostTweet(1, 101) // User 2 posts a new tweet with id = 20.
	twitter.PostTweet(1, 13)  // User 2 posts a new tweet with id = 20.
	twitter.PostTweet(1, 10)  // User 2 posts a new tweet with id = 20.
	twitter.PostTweet(1, 94)  // User 2 posts a new tweet with id = 20.
	twitter.PostTweet(1, 505) // User 2 posts a new tweet with id = 20.
	twitter.PostTweet(1, 333) // User 2 posts a new tweet with id = 20.
	twitter.PostTweet(1, 22)  // User 2 posts a new tweet with id = 20.
	fmt.Printf(
		"Feed for 1 before following %+v\n", twitter.GetNewsFeed(1),
	) // User 1's news feed should only contain their own tweets -> [10].
}

func main() {
	test2()
}

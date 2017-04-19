package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

/*func getUserByScreenName(ScreenName string) *twitter.User{


}*/
func getUserToFollowFromUser(client *twitter.Client, fromUser string) string {
	friends, httpResp, err := client.Friends.List(&twitter.FriendListParams{
		ScreenName: fromUser,
		Count:      1,
	})
	if err != nil {
		fmt.Println(err)
	}
	if httpResp.Status != "200 OK" {
		c.Red(httpResp.Status)
	}
	//fmt.Println(friends.Users)
	c.Green(friends.Users[0].ScreenName)
	return friends.Users[0].ScreenName
}

func followUser(client *twitter.Client, screenName string) {
	_, _, _ = client.Friendships.Create(&twitter.FriendshipCreateParams{
		ScreenName: screenName,
	})
}

func followRandom(client *twitter.Client, nFollow int, screenName string) {
	fmt.Println("Starting to follow " + strconv.Itoa(nFollow) + " users")
	//screenName := firstScreenName
	for i := 0; i < nFollow; i++ {
		userToFollow := getUserToFollowFromUser(client, screenName)
		followUser(client, userToFollow)
		screenName = userToFollow
		//wait to avoid the twitter api limitation
		fmt.Println("waiting 1 min to avoid twitter api limitation")
		fmt.Println(time.Now().Local())
		time.Sleep(1 * time.Minute)
	}

}

func optionFollowRandom(client *twitter.Client) {
	c.Red("how many accounts to follow?")
	newcommand := bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	fmt.Print("Number of users to follow: ")
	c.Purple(answer)
	nFollow, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("incorrect entry, need a positive number")
	}

	fmt.Print("entry @username of a user, to get a 1st user to follow, that will be a user that the 1st user is following, and the 2nd user will be a user that the 3rd user is following): @")
	newcommand = bufio.NewReader(os.Stdin)
	firstScreenName, _ := newcommand.ReadString('\n')
	firstScreenName = strings.TrimSpace(firstScreenName)
	fmt.Print("first user to follow: @")
	c.Purple(firstScreenName)

	c.Red("Are you sure? [y/n]")
	newcommand = bufio.NewReader(os.Stdin)
	answer, _ = newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		followRandom(client, nFollow, firstScreenName)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}

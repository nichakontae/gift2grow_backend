package notification

import (
	"context"
	"fmt"
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/response"
	"log"
	"strconv"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

func NotifyUser(c *fiber.Ctx) error {

	//initialize firebase app
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	//initialize firebase messaging client
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	campaignId := c.Query("campaignId")
	if campaignId == "" {
		return &response.GenericError{
			Message: "campaignId is required",
			Err:     nil,
		}
	}

	//get userIds
	var userCampaign []model.DonateHistory

	if result := mysql.Gorm.Where("campaign_id = ?", campaignId).Find(&userCampaign); result.Error != nil {
		return &response.GenericError{
			Message: "Error when query user campaign",
			Err:     result.Error,
		}
	}

	if len(userCampaign) == 0 {
		return &response.GenericError{
			Message: "No user campaign found",
			Err:     nil,
		}
	}

	userIds := make([]string, len(userCampaign))
	for i, v := range userCampaign {
		if v.UserId != nil {
			userIds[i] = *v.UserId
		}
	}

	//print("userIds: ", userIds)

	//post notification to database
	campaignID, err := strconv.ParseUint(campaignId, 10, 64)
	if err != nil {
		return &response.GenericError{
			Message: "Invalid campaignId format",
			Err:     err,
		}
	}

	NotiObject := &model.NotiObject{
		CampaignId: &campaignID,
		UserNoti:   []*model.UserNoti{},
	}

	isRead := false
	for _, userID := range userIds {
		currentUserID := userID // Create a separate variable to store the current userID
		userNoti := &model.UserNoti{
			UserId: &currentUserID,
			IsRead: &isRead,
		}
		NotiObject.UserNoti = append(NotiObject.UserNoti, userNoti)
	}

	if result := mysql.Gorm.Create(NotiObject); result.Error != nil {
		return &response.GenericError{
			Message: "Error when create notification object",
			Err:     result.Error,
		}
	}

	print("NotiObject: ", NotiObject)
	print("NotiObject.UserNoti: ", NotiObject.UserNoti)

	//push notification to users
	for _, v := range userIds {

		//get tokens
		var userTokens []model.UserToken

		if result := mysql.Gorm.Where("user_id = ?", v).Find(&userTokens); result.Error != nil {
			return &response.GenericError{
				Message: "Error getting user token",
				Err:     result.Error,
			}
		}

		// print(
		// 	"userId: ", v,
		// 	"tokens: ", userTokens,
		// )
		if len(userTokens) == 0 {
			return &response.GenericError{
				Message: "No user token found",
				Err:     nil,
			}
		}

		tokens := make([]string, len(userTokens))
		for i, userToken := range userTokens {
			if userToken.Token != nil {
				tokens[i] = *userToken.Token
			}
		}

		//create the message
		message := &messaging.MulticastMessage{
			Notification: &messaging.Notification{
				Title: "Gift2Grow",
				Body:  "The campaign you are following has been updated the evidence image",
			},
			Tokens: tokens,
			Data: map[string]string{
				"CampaignId": campaignId,
			},
		}

		//send the message
		response, err := client.SendMulticast(context.Background(), message)
		if err != nil {
			log.Fatalln(err)
		}

		//print the message ID token
		fmt.Println("Successfully sent message:", response)
	}

	return c.JSON("success")

}

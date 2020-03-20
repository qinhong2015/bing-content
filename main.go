package main

import (
	"FeedImport/Service"
	"FeedImport/db"
	"bingContent"
	_ "bingContent"
	"fmt"
)

func main() {
	fmt.Print("xiao bao bao is so cute")

	env := "staging"
	accountId := "sando_5de81a0034378"
	config, err := Service.GetS3Service().GetConfig(env)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dbs := config["dbs"].(map[string]interface{})["index_readonly"]
	indexPdo := db.CreateInstance(dbs, "sando_index")
	Service.GetAccountService().SetAccount(indexPdo, accountId, config)

	var clientIdString, clientSecretString string

	clientId := config["bing"].(map[string]interface{})["client_id"]
	clientSecret := config["bing"].(map[string]interface{})["client_secret"]
	//developerToken := config["bing"].(map[string]string)["developer_token"]
	clientIdString = fmt.Sprintf("%s", clientId)
	clientSecretString = fmt.Sprintf("%s", clientSecret)

	account := Service.GetAccountService().Account
	account.LoadSettings()
	token := account.GetSettings("bing_refresh_token")
	merchantId := account.GetSettings("bing_merchant_id")
	customerId := account.GetSettings("bing_customer_id")
	adsAccountId := account.GetSettings("bing_ads_id")
	isMicrosoftIdentityPlatform := account.GetSettings("bing_is_microsoft_identity_platform")
	if isMicrosoftIdentityPlatform == nil {
		isMicrosoftIdentityPlatform = false
	}

	oAuthWithAuthorizationCode := bingContent.NewOAuthWithAuthorizationCode()
	oAuthWithAuthorizationCode.OAuthAuthorization.WithClientSecret(clientSecretString)
	oAuthWithAuthorizationCode.OAuthAuthorization.WithClientId(clientIdString)
	if isMicrosoftIdentityPlatform == true {
		oAuthWithAuthorizationCode.OAuthAuthorization.WithRequireLiveConnect(true)
	}

	tokenString := fmt.Sprintf("%s", token)
	accesstoken := oAuthWithAuthorizationCode.RequestOAuthTokensByRefreshToken(tokenString)

	fmt.Printf("%s", merchantId)
	fmt.Printf("%s", accesstoken)

	fmt.Printf("%s", token)

	fmt.Printf("%s", customerId)
	fmt.Printf("%s", adsAccountId)
	fmt.Printf("%s", isMicrosoftIdentityPlatform)
}

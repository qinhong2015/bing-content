package main

import (
	"FeedImport/Service"
	"FeedImport/db"
	"fmt"
	"log"
)

func main() {
	env := "staging"
	accountId := "sando_5e1619fc29ba0"
	config, err := Service.GetS3Service().GetConfig(env)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dbs := config["dbs"].(map[string]interface{})["index_readonly"]
	indexPdo := db.CreateInstance(dbs, "sando_index")
	Service.GetAccountService().SetAccount(indexPdo, accountId, config)
	rows, err := indexPdo.Query("SELECT `a`.`name`, `a`.`account_id`, `a`.`source`, `a`.`mysql_host`, `a`.`timezone`, `a`.`source`, CASE WHEN ap.permissions LIKE '%\"product_modify_advance_modifier\": true%' THEN 1 ELSE 0 END AS `advance_modifier_enabled`, CASE WHEN ap.permissions LIKE '%\"product_modify_promotion_id\": true%' THEN 1 ELSE 0 END AS `modify_promotion_id_enabled`, CASE WHEN ap.permissions LIKE '%\"product_modify_google_express\": true%' THEN 1 ELSE 0 END AS `product_modify_google_express_enabled` FROM `accounts` AS `a` INNER JOIN `account_permissions` AS `ap` ON `ap`.`account_id` = `a`.`account_id` WHERE `a`.`account_id` = ?;", accountId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()	//indexDbs := dbs["index_readonly"]

	rows.Scan()
	fmt.Println(dbs)
	fmt.Println(indexPdo)
}

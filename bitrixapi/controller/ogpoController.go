package controller

import (
	"bitrixapi/database"
	"bitrixapi/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Test = func(res http.ResponseWriter, req *http.Request) {
	test := map[string]interface{}{"status": true, "val": "zbs"}
	json.NewEncoder(res).Encode(test)
}

var InsertCust = func(res http.ResponseWriter, req *http.Request) {
	ogpo := &models.Ogpo{}
	err := json.NewDecoder(req.Body).Decode(ogpo)
	log.Println(err)
	if err != nil {
		log.Println(err)
		resp := map[string]interface{}{"status": false, "message": "Неправильный боди"}
		json.NewEncoder(res).Encode(resp)
		return
	}

	InsertData(*ogpo)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(res).Encode(map[string]interface{}{"status": true, "message": "row inserted successfully"})
}

func InsertData(ogpo models.Ogpo) {
	tx := database.MYSQLDBGORM.Begin()

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	_ = tx.Exec(`INSERT INTO b_iblock_element 
			(TIMESTAMP_X, MODIFIED_BY, 
			DATE_CREATE, CREATED_BY, 
			IBLOCK_ID, ACTIVE, 
			SORT, NAME, 
			PREVIEW_TEXT_TYPE, DETAIL_TEXT_TYPE, 
			SEARCHABLE_CONTENT, WF_STATUS_ID, 
			IN_SECTIONS)
			VALUES (NOW(), 4, NOW(), 4, 25, 'Y', 500, ?, 'text', 'text', ?, 1, 'N')`, ogpo.Name, ogpo.Name)
	//if err != nil {
	//	fmt.Println(err.Error)
	//	tx.Rollback()
	//	//return err.Error
	//}

	_ = tx.Exec(`UPDATE b_iblock_element set XML_ID = (select id from b_iblock_element
							where SEARCHABLE_CONTENT = ?
							order by id desc
							limit 1)
							where id = (select id from b_iblock_element
							where SEARCHABLE_CONTENT = ?
							order by id desc
							limit 1)`, ogpo.Name, ogpo.Name)


	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (142, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.LastName)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (143, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.FirstName)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (144, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.FatherName)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (145, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.City)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (146, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.PhoneNumber)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (147, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.Email)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (148, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.IIN)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (149, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.ExpAge)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (150, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.VehicleType)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (151, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.RegistrationVehLoc)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (152, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.UsageDuration)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (161, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.TotalPolicyPrice)

	_ = tx.Exec(`INSERT INTO b_iblock_element_property (IBLOCK_PROPERTY_ID, IBLOCK_ELEMENT_ID, VALUE, VALUE_TYPE)
					  VALUES (155, (select id from b_iblock_element
					  where SEARCHABLE_CONTENT = ?
					  order by id desc
					  limit 1), ?, 'text');`, ogpo.Name, ogpo.Language)

	tx.Commit()
}

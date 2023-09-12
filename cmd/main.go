package main

import (
	"fmt"
	"log"

	"github.com/hellosumitg/realtimedb/rtdb"
)

func main() {
	db, err := realtimedb.New()
	if err != nil {
		log.Fatal(err)
	}	
	
	// data := map[string]any{} // int, string, []byte, float, email datatypes....
	user := map[string]string{
		"name": "Sumit",
		"age":  "30",
	} // int, string, []byte, float, ....

	id, err := db.Insert("users", user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", id)

	// db.Update(func(tx *bbolt.Tx) error {
	// 	bucket, err := tx.CreateBucket([]byte("users"))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	
		

	// 	return nil
	// })

	// userData := make(map[string]string)
	// if err := db.View(func(tx *bbolt.Tx) error {
	// 	bucket := tx.Bucket([]byte("users"))
	// 	if bucket == nil {
	// 		return fmt.Errorf("bucket (%s) not found", "users")
	// 	}

	// 	bucket.ForEach(func(k, v []byte) error{
	// 		userData[string(k)] = string(v)
	// 		return nil
	// 	})

	// 	return nil
	// }); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(user)
}

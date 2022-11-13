package main

/*
go get -u github.com/go-sql-driver/mysql
*/
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id     int
	Name   string
	Gender string
	Status string
}

func main() {
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	var connectionString = "root:nephilim17@tcp(127.0.0.1:3306)/tugas_ddl?parseTime=true"
	db, err := sql.Open("mysql", connectionString) // membuka koneksi ke daatabase
	if err != nil {                                // pengecekan error yang terjadi ketika proses open connection
		log.Fatal("error open connection", err.Error())
	}

	errPing := db.Ping() // mengecek apakah apliasi masih terkoneksi ke database
	if errPing != nil {  //handling error ketika gagal konek ke db
		log.Fatal("error connect to db ", errPing.Error())
	} else {
		fmt.Println("koneksi berhasil")
	}

	defer db.Close() // menutup koneksi

	//buat mekanisme menu
	fmt.Println("MENU:\n1. Baca data\n2. Tambah data\n3. Update data\n4. Delete data\n5. Baca data by ID")
	fmt.Println("Masukkan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			result, errSelect := db.Query("SELECT id, name, gender, status FROM users") // proses menjalankana query SQL
			if errSelect != nil {                                                       //handling error saat proses menjalankan query
				log.Fatal("error select ", errSelect.Error())
			}

			var dataUser []User
			for result.Next() { // membaca tiap baris/row dari hasil query
				var userrow User                                                                     // penampung tiap baris data dari db                                                                                                     // membuat variabel penampung
				errScan := result.Scan(&userrow.Id, &userrow.Name, &userrow.Gender, &userrow.Status) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
				if errScan != nil {                                                                  // handling ketika ada error pada saat proses scannign
					log.Fatal("error scan ", errScan.Error())
				}
				// fmt.Printf("id: %s, nama: %s, email: %s\n", userrow.Id, userrow.Nama, userrow.Email) // menampilkan data hasil pembacaan dari db
				dataUser = append(dataUser, userrow)
			}
			// fmt.Println("data all", dataUser)
			for _, value := range dataUser { // membaca seluruh data user yang telah ditampung di variable slice
				fmt.Printf("id: %d, name: %s, gender: %s, status: %s\n", value.Id, value.Name, value.Gender, value.Status)
			}
		}
	case 2:
		{
			newUser := User{}
			fmt.Println("masukkan name user")
			fmt.Scanln(&newUser.Name)
			fmt.Println("masukkan gender user")
			fmt.Scanln(&newUser.Gender)
			fmt.Println("masukkan status user")
			fmt.Scanln(&newUser.Status)

			statement, errPrepare := db.Prepare(`INSERT INTO users (name, gender, status) VALUES (?, ?, ?)`)
			if errPrepare != nil {
				log.Fatal("error prepare insert", errPrepare.Error())
			}

			result, errExec := statement.Exec(newUser.Name, newUser.Gender, newUser.Status)
			if errExec != nil {
				log.Fatal("error exec insert", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("Insert berhasil")
				} else {
					fmt.Println("Insert gagal")
				}
			}
		}
	case 3:
		{

			updateUser := User{}
			fmt.Println("masukkan id user yang akan di update")
			fmt.Scanln(&updateUser.Id)
			fmt.Println("masukkan name user")
			fmt.Scanln(&updateUser.Name)
			fmt.Println("masukkan gender user (M/F)")
			fmt.Scanln(&updateUser.Gender)
			fmt.Println("masukkan status user (active/inactive")
			fmt.Scanln(&updateUser.Status)

			statement, errPrepare := db.Prepare(`UPDATE users set name = ?, gender = ?, status = ? where id = ?`)
			if errPrepare != nil {
				log.Fatal("error prepare update", errPrepare.Error())
			}cd ..errPrepare

			result, errExec := statement.Exec(updateUser.Name, updateUser.Gender, updateUser.Status, updateUser.Id)
			if errExec != nil {
				log.Fatal("error exec update", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("update berhasil")
				} else {
					fmt.Println("update gagal")
				}
			}

		}

	case 4:
		{
			// fmt.Println("delete")
			deleteUser := User{}
			fmt.Println("masukkan id user yang akan di DELETE")
			fmt.Scanln(&deleteUser.Id)

			statement, errPrepare := db.Prepare(`DELETE from users where id = ?`)
			if errPrepare != nil {
				log.Fatal("error prepare delete", errPrepare.Error())
			}

			result, errExec := statement.Exec(deleteUser.Id)
			if errExec != nil {
				log.Fatal("error exec delete", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("delete berhasil")
				} else {
					fmt.Println("delete gagal")
				}
			}

		}

	case 5:
		{
			// fmt.Println("baca data by id")
			bacaUser := User{}
			fmt.Println("masukkan id user")
			fmt.Scanln(&bacaUser.Id)

			results := db.QueryRow("SELECT id, name, gender, status from users where id = ?", &bacaUser.Id)
			var dataUser User
			err := results.Scan(&dataUser.Id, &dataUser.Name, &dataUser.Gender, &dataUser.Status)

			if err != nil {
				log.Fatal("error select ", err.Error())
			}
			fmt.Printf("id: %d, name: %s, gender: %s, status: %s\n", dataUser.Id, dataUser.Name, dataUser.Gender, dataUser.Status)
			//----------------

		}

	}

}

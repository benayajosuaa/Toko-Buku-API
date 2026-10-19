package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"os"
)



type Buku struct {
	ID int `json:"id"`
	Judul string `json:"judul"`
	Penulis string `json:"penulis"`
	Harga int `json:"harga"`
	Stok int `json:"stok"`
}

type DataBuku struct {
	DaftarBuku  []Buku `json:"data_buku"`
}

// CREATE

// GET
func getAllBook(w http.ResponseWriter, r *http.Request){
	// validasi utk menggunakan method get
	if r.Method != http.MethodGet{
		w.Header().Set("Allow", "GET")
		http.Error(w, "Gunakan Method GET", http.StatusMethodNotAllowed)
		return
	}

	// baca isi file
	isiFile, err := os.ReadFile("data.json")
	if err != nil {
		log.Println("Gagal membaca file", err)
		http.Error(w, "Gagal membaca data buku", http.StatusInternalServerError)
		return
	}

	// siapkan variabel untuk menampung data JSON
	var data DataBuku

	// Ubah isi JSON menjadi data GO
	err = json.Unmarshal(isiFile, &data)
	if err != nil {
		log.Println("Gagal memproses JSON:", err)
		http.Error(w, "Gagal memproses data buku", http.StatusInternalServerError)
		return	
	}

	if data.DaftarBuku == nil {
		data.DaftarBuku = []Buku{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data.DaftarBuku)
	if err != nil {
		log.Println("Gagal mengirim respons: ", err)
	}
	
}



// GET by ID

// UPDATE

// DELETE

func main(){
	// deklarasi routing atau membuat router
	mux := http.NewServeMux()

	mux.HandleFunc("/buku", getAllBook)

	// membuat Endpoint API
	mux.HandleFunc( "/", func(w http.ResponseWriter, r *http.Request){
		// Status 
		w.WriteHeader(http.StatusOK)
		// Pesan atau Response API
		fmt.Fprintln(w, "API Toko Bukunya sukses dibuat")
		fmt.Fprintln(w, "Versi 0.0.1 - 18.10.2026")
		fmt.Fprintln(w, "dibuat oleh Michael")
		return
	})
	// Print Log di terminal
	log.Println("Server berhasil jalan di http://localhost:8080")

	// Menjalankan HTTP Server, 
	// kemudian menampilkan error dan juga menghentikan program jika ListenandServe nya return Eror
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}


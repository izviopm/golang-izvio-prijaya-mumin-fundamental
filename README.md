# Project-1

Project ini merupakan aplikasi berbasis console yang dibuat menggunakan bahasa Go, aplikasi ini bertujuan untuk Mengelola Data buku, Data data buku disimpan kedalam local variable menggunakan slice sehingga ketika aplikasi selesai digunakan data akan kembali ke data default

## Fitur

- Melihat Semua Data Buku
- Menambahkan Data Buku
- Menampilkan Data Berdasarkan ID
- Mengubah Salah Satu Data Menggunakan ID
- Menghapus Salah Satu Data Berdasarkan ID

## Flowchart

Berikut adalah flowchart dari setiap fitur yang ada dalam project-1

1. Melihat Semua Data
   <br>
   <br>
   <img src="/project-1/assets/image/1. Show All.png" alt="Melihat Semua Data" style="width:30%;">
   
2. Menambahkan Data Buku
   <br>
   <br>
   <img src="/project-1/assets/image/2. Add Data Buku.png" alt="Add Data" style="width:26%;">

3. Menampilkan Data Berdasarkan ID
   <br>
   <br>
   <img src="/project-1/assets/image/3. Filter Data.png" alt="Filter Data" style="width:52%;">

4. Mengubah Salah Satu Data Menggunakan ID
   <br>
   <br>
   <img src="/project-1/assets/image/4. Update Data.png" alt="Update Data" style="width:52%;">

5. Menghapus Salah Satu Data Berdasarkan ID
   <br>
   <br>
   <img src="/project-1/assets/image/5. Delete Data.png" alt="Delete Data" style="width:52%;">

## Time Complexity

1. Inisialisasi data awal: O(1) Operasi ini memiliki kompleksitas waktu tetap karena hanya melibatkan inisialisasi variabel dengan nilai awal.
2. Loop utama: O(N) - Loop ini akan berjalan sebanyak N kali, di mana N adalah jumlah iterasi yang dilakukan oleh pengguna sebelum memasukkan perintah "exit" atau "6". Jumlah iterasi ini bergantung pada interaksi pengguna dengan program.
3. Pemeriksaan perintah: O(1) - Pemeriksaan perintah dilakukan dengan menggunakan pernyataan if-else. Operasi ini memiliki kompleksitas waktu tetap karena hanya melibatkan pemeriksaan terhadap nilai yang sudah diketahui.
4. Menampilkan data dalam bentuk tabel: O(N) - Operasi ini melibatkan iterasi sebanyak N kali untuk menampilkan data dalam bentuk tabel. Jumlah iterasi ini sama dengan jumlah data yang ada.
5. Menambahkan data baru: O(1) - Operasi ini memiliki kompleksitas waktu tetap karena hanya melibatkan penambahan elemen ke dalam slice data.
6. Mencari data dengan filter: O(N) - Operasi ini melibatkan iterasi sebanyak N kali untuk mencari data yang sesuai dengan filter. Jumlah iterasi ini sama dengan jumlah data yang ada.
7. Mengubah data: O(N) - Operasi ini melibatkan iterasi sebanyak N kali untuk mencari data yang sesuai dengan filter dan mengubahnya. Jumlah iterasi ini sama dengan jumlah data yang ada.
8. Menghapus data: O(N) - Operasi ini melibatkan iterasi sebanyak N kali untuk mencari data yang sesuai dengan filter dan menghapusnya. Jumlah iterasi ini sama dengan jumlah data yang ada.

Berdasarkan analisis tersebut, kompleksitas waktu program secara keseluruhan dapat dinyatakan sebagai O(N), di mana N adalah jumlah data yang ada. Program akan berjalan lebih lama jika jumlah data semakin besar.

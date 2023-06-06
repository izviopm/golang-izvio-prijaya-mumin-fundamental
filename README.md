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
   <img src="/project-1/assets/image/2. Show All.png" alt="Melihat Semua Data" width="400" height="600">

2. Menambahkan Data Buku
   ![Menambahkan Data Buku](project-1/assets/image/2.%20Add%20Data%20Buku.png)

3. Menampilkan Data Berdasarkan ID
   ![Menampilkan Data Berdasarkan ID](project-1/assets/image/3.%20Filter%20Data.png)

4. Mengubah Salah Satu Data Menggunakan ID
   ![Mengubah Salah Satu Data Menggunakan ID](project-1/assets/image/4.%20Update%20Data.png)

5. Menghapus Salah Satu Data Berdasarkan ID
   ![Menghapus Salah Satu Data Berdasarkan ID](project-1/assets/image/5.%20Delete%20Data.png)

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

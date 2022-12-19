# Go Backend Template

Ini merupakan program template untuk Backend Go menggunakan REST API.

## Requirements

Ini adalah requirement untuk menjalankan template ini:

1. Go v1.19.3
2. Docker (Optional)
3. Postgresql (Optional, when using docker)
4. GNU Make 4.3 (Optional, when using docker)
5. Minio (Optional, when using docker)
6. Google wire v0.5.0 (https://github.com/google/wire)
7. Air command line (https://github.com/cosmtrek/air)

## Cara menjalankan

Untuk menjalankan server ini, gunakan:

```
make watch
```

Atau, jika menggunakan docker, kamu dapat gunakan:

```
docker compose up
```

## Alur Program

Program ini akan dimulai dari modul app. Saat ada request masuk, berikut alurnya:

1. Program akan memanggil modul `middleware` untuk menjalankan seluruh middleware yang terdaftarkan.
2. Program akan memanggil module `routes` untuk menyocokan rute yang sesuai.
3. Saat menemukan rute yang sesuai, modul `routes` akan memanggil modul `controller`.
4. Modul controller akan melakukan parsing data dan validasi input pengguna lalu akan melempar hasil parsing pada modul `service`.
5. Pada modul `service`, proses bisnis akan dilakukan. Modul ini bisa saja memanggil modul `repository` untuk mengambil data, modul `lib.storage` untuk melakukan operasi dengan S3 Bucket, dan lainnya.
6. Hasil operasi `service` akan dikembalikan pada `controller`. Controller akan melakukan parsing data dan mengirimkan hasilnya pada pengguna.

## Struktur Folder

Berikut ini adalah struktur proyek ini:

### Folder `app`

Folder `app` digunakan untuk menyimpan kelas yang berkaitan dengan sistem internal server. Terdapat beberapa file penting diantarnaya adalah sebagai berikut:

- File `di.go` digunakan untuk menambahkan referensi pada dependency injection. Injector yang kami gunakan adalah Google wire.
- File `wire_gen.go` merupakan file hasil generate dari google wire. File ini tidak boleh diubah sama sekali.

### Folder `config`

Folder `config` merupakan folder yang digunakan untuk menghubungkan environment variable dengan sistem. Setiap field konfigurasi perlu memiliki `types` yang berbeda apabila ingin dilakukan inject pada kelas lain.

Setiap kali pembuatan type baru, pastikan untuk membuat implementasi method `UnmarshalText` agar env dapat diconvert menjadi types tersebut.

Setiap field yang akan ditambahkan pada program ini perlu didaftarkan pada struct `config`. Lalu, agar field tersebut dapat diakses melalui dependency injection, field tersebut perlu ditambahkan pada `app/di.go`.

### Folder `controller`

Folder `controller` ditujukan untuk menyimpan semua controller yang dibutuhkan. Controller harus dikelompokan sesuai dengan fungsinya. Controller yang memiliki kategori fungsi yang berbeda perlu dibuat modul terpisah untuk kerapihan.

Setiap modul controller perlu memiliki interface yang menyatakan daftar fungsi controller yang tersedia. Fungsi yang terdaftar di controller haruslah bertipe `http.HandlerFunc`.

Pada `controller` terdapat beberapa proses yang perlu dilakukan:

1. Melakukan parsing body data serta data-data yang dibutuhkan oleh `service`.
2. Melakukan validasi input.
3. Menambahkan dokumentasi mengenai endpoint.

### Folder `docs`

Folder ini menyimpan dokumentasi yang digenerate dari swag.

### Folder `embed`

Folder `embed` digunakan untuk menyimpan file-file yang akan dicompile bersama program.

### Folder `lib`

Folder `lib` ditujukan untuk menyimpan juga kode program yang berkaitan dengan sistem dan juga kode program yang tidak berkaitan dengan folder lainnya.

### Folder `middleware`

Folder middleware menyimpan segala middleware yang tersedia pada program.

### Folder `model`

Folder ini digunakan untuk menyimpan struktur data dari program. Terdapat dua buah subfolder yang ada pada folder ini, yaitu:

- Folder `domain` digunakan untuk menyimpan struktur data basis data.
- Folder `web` digunakan untuk menyimpan struktur data respons.

### Folder `repository`

Folder ini digunakan untuk menyimpan handler menuju basis data. Repository hanya boleh menghubungkan satu buah tabel saja, bila memerlukan pemrosesan lebih dari dua tabel, lakukan pada `service`.

### Folder `routes`

Folder ini berisi routes yang digunakan pada API ini. Semua struct yang ada pada kelas ini harus memiliki skema sebagaimana interface `BaseChiRoute`.

### Folder `service`

Folder ini untuk menyimpan segala kode program yang berkaitan dengan pemrosesan request.

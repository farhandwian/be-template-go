### Register user `POST /register` by ADMIN
Admin akan mendaftarkan User yang baru join dengan form `POST /register` dengan menyertakan informasi
* Name
* Email
* PhoneNumber

Aktivitas ini perlu dilanjut ke API `POST /activation/request`

---

### Send email activation request `POST /activation/request` by FRONTEND / ADMIN 
Setelah API `POST /register` selesai, Frontend akan memanggil API `POST /activation/request`. 

Kedua API ini sengaja dipisah untuk mengantisipasi kegagalan pengiriman email. Frontend bisa menyertakan tombol `POST /activation/request` untuk kembali melakukan trigger kirim email. 

Tombol ini harusnya di disabled (atau bahkan di hide) jika user sudah teraktivasi. 

User sudah teraktivasi ditandai dengan field `email_verified_at` sudah berisi suatu nilai timestamp. 

Aktivitas ini akan mentrigger pengiriman email ke user yang bersangkutan

---

### Submit email activation `POST /activation/submit` by ANONYMOUS
User yang belum login, akan membuka email lalu akan melihat sebuah link 
`https://aplikasi.com/user-activation-form?token=abc123`
yang jika diklik akan membuka sebuah halaman frontend yang berisi form input. User harus mengisikan
* Password
* Pin (4 digit)
* ActivationToken (sebagai hidden field yang dari token)

Lalu submit form ini ke `POST /activation/submit`

---

### Initiate user login `POST /login` by ANONYMOUS
User yang belum login masuk ke halaman login lalu mengisikan
* Email
* Password

Lalu submit form ini ke `POST /login`

Aktivitas ini akan mentrigger pengiriman OTP ke whatsapp

Jika berhasil, halaman frontend akan redirect ke pengisian OTP

---

### Submit OTP for login `POST /login/otp` by ANONYMOUS
User lalu akan mengisi OTP yang diterima dari whatsapp kedalam form ini.
* OTP (6 digit)
* Email (hidden field)

Lalu submit form ini ke `POST /login/otp`

Akan mereturn `AccessToken` dan `RefreshToken`

Jika berhasil, halaman frontend akan redirect ke dashboard

---

### Get all users `GET /user` by ADMIN
Admin bisa melihat seluruh user dengan menggunakan API ini
Untuk tiap user, Admin bisa melakukan 
* Trigger aktivasi email manual
* Trigger reset password
* Melihat, memberikan dan mencabut hak akses
* Melakukan suspend user (not implemented yet)
* Melakukan update data user (not implemented yet)

Beberapa filter yang bisa digunakan :
* Email
* NameLike
* UserID
* PhoneNumber
* Page dan Size untuk paging

---

### Initiate password reset `POST /password/reset/request` by ADMIN
Admin memerlukan `GET /user` untuk mencari user yang akan direset passwordnya.

Aktivitas ini akan mentrigger pengiriman email ke user yang bersangkutan

---

### Submit password reset `POST /password/reset/submit` by ANONYMOUS
User yang kehilangan password akan membuka email dan mendapati link berupa
`https://aplikasi.com/user-password-reset?token=abc123`
yang jika diklik akan membuka sebuah halaman frontend yang berisi form input. User harus mengisikan
* NewPassword (dibuat 2 field sebagai konfirmasi)
* PasswordResetToken (as hidden field)

Lalu akan memanggil field `POST /password/reset/submit`

---







** strategi deployment
* big-bang/replace sebagai contoh ketika ada 3 servir maka ketiga server akan dimatikan dan diupgrade( kelebihan -> mudah diimplementasikan karena tinggal melakukan replace dan perubahan kepada sistem langsung secara instant, kekurangan -> beresiko dan rata2 downtime cukup lama)
* Rollout ketika memiliki 3 server maka server pertama dulu dijadikan new version, ketika sudah selesai maka lanjut ke server 2 dan seterusnya(kelebihan -> lebih aman dan lessd downtime, kekurangan -> akan ada 2 versi apk yang berjalan bersamaan, untuk deployment dan rollback lebih lama dan tidak ada kontrol request)
* blue/green menyiapkan server baru untuk versi terbaru sehingga ketika telah selesai mealkukan set up maka aplikasi server langsung dapat(kelebihan -> perubahan sangat cepat, tidak ada issue beda versi, kekurangan ->resourch lebih banyak)
* canary hampir sama dengan blue/green namun pada canary ketika melakukan switch secara perlahan dan tidak instant

* deploy gcp :

1. buat vm instance lalu install yang diperlukan (compute engine)
2. buat sql di google cloud lalu buat database sesuai config
3. ubah network pada connections di sql lalu add network dengan ip 0.0.0.0/0 untuk mengizinkan setiap ip mengakses
4. ubah pada firewall policies gcp untuk protocols and ports allow all
5. masuk kedalam vm instance lalu pindahkan project kedalam 1 folder didalam instance
6. jalankan dockerfile yang terdapat didalam project dengan build image lalu jalankan pada container
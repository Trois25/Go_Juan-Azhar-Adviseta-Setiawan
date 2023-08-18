#version control and branch management(git)
*versioning adalah pengaturan versi/pelacakan perubahan pada code kita

* Version Control system yang akan digunakan adalah GIT yang mana dapat melakukan kolaborasi dengan software enginer lainnya

** GIT repository
* git akan melakukan tracking pada file
* git dapat melakuukan undo yang disebut commit
* pada git code akan disimpan didalam repository yang mana merupakan server yang menyimpan code

**Melakukan remote agar local terhubung dengan server dengan cara :
* git init
* git remote add origin "link repository"

**beberapa command pada github github :
* git add . ,untuk mengupload semua perubahan ke stagging area
* git commit -m "pesan yang dikirim ke github", untuk memastikan file project benar untuk dilakukan push
* git push -u origin main , untuk melakukan upload file kedalam github
* git diff , untuk melihat perubahan apa saja yang terjadi pada code
* git stash ,untuk memindahkan perubahan yang telah dibuat pada local ke stash area, stash ini dapat dikembalikan dengan cara git stach apply
* git log , digunakan untuk melihat log dari github
* git reset commit_id(dapat dilihat menggunakan git log -oneline) --soft(--soft/--hard) digunakan untuk mengembalikan ke stage sebelumnya. soft akan mengembalikan ke commid_id yang dipilih dan menyimpan file sebelumnya ke staged area sedangkan hard akan benar2 menghapusnya.
* git fetch digunakan untuk mengambil data didalam server dan melihat branch mana yang diupdate, command ini akan mengambil update namun tidak akan masuk kedalam local dan hanya menampilkan perubahan yang terjadi di server.
* git pull, gabungan git fetch dan git master. fungsinya mengambil data perubahan dan memasukkannya kedalam local	

**file .gitignore, didalam file ini dituliskan file/folder apapun yang tidak ingin di push kedalam github




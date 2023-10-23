** library go untuk machine learning
* GoLearn,merupakan library machine learning yang batteries included ( tinggal menggunakan) 
* Gorgonia, bekerja dengan membuat computation graph yang terbatas mengerjakan kedalam matematican function
* gonum, package untuk mengolah data numerical dan algoritma saint menjadi produktif, memiliki perforam tinggi dan scalable
* gomind merupakan neural network library yang hanya support single hidden layer

** untuk mengembangkan model aplikasi berbasis AI ada cara yang lebih mudah yakni menggunakan teknologi A.I as a Service (AIaaS) yakni menggunakan model yang sudah jadi

** Tipe-tipe AIaaS
* Bots/Chatbots, membuat AI berbasis customer service. ct : amazon lex (AWS), Azure bot service (microsoft azure), DialogFlow(GCP)
* APIs, mengintegrasikan AI milik vendor dengan aplikasi sendiri. ct : amazon rekognition, amazon comprehend, azure cognitive service, azure speech services, google cloud vision, cloud natural language
* Machine learning, build & deploy machine learning model lebih mudah. ct : amazon sagemaker, azure machine learning, google cloud AI

** Keuntungan AIaaS
* Pengurangan cost
* Low-code
* Pross deployment cepat
* Flexibility
* Usability
* Scalability
* Customization

** Kelemahan AIaaS
* Cost yang berlebih dalam jangka panjang
* Privasi data
* Keamanan
* Transparansi
* Vendor lock-in

** OpenAI API
* Target : mendesain prompt yang tepat (prompt engineering) dan membuat aplikasi berbasis A.I. menggunakan API OpenAI

** Mendapatkan API key untuk mengakses OpenAI API
* langkah 1 : buka situs openai
* langkah 2 : buat akun openai untuk mendapatkan api key
* langkah 3 : cari bagian introduction dan masuk ke API keys
* langkah 4 : pilih API keys dan create new secret key
* langkah 5 : ikuti petunjuk untuk mengkonfigurasi api key sesuai kebutuhan
* langkah 6 : copy secrete key

*** Prinsip menggunakan prompt yang baik

** gunakan prompt yang jelas & spesifik. 
* gunakan delimiters (```,"", <>,<tag></tag>) untuk menandakan bagian mana yang menjadi input
* tuliskan struktur output yang diinginkan * minta model untuk memeriksa apakah input sesuai
* Few-shot prompting

*** Berikan "waktu berpikir" untuk menghindari solusi yang salah
** Menulis langkah-langkah yang dibutuhkan untuk menyelesaikan tugas dan output yang diinginkan
** menginstruksikan model untuk menuliskan solusinya sendiri sebelum masuk ke kesimpulan 
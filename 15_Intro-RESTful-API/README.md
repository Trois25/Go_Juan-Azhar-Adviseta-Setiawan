** Restful API

* penggunaan restful API digunakan untuk membuat suatu aplikasi dapat memenuhi kebutuhan front end yang bertujuan untuk multi platform.

** Kegunaan dan pengimplementasian dari restful API :

* Frontend Backend Integration, cara kerja dari API adalah terjadinya request dari frontend kepada backend. setelah dilakukannya request maka BE akan merespon sesuai dengan request yang diberikan oleh FE.

* Service to Service, yakni dapat mengintegrasikan antara BE ke BE.

** API Tools (Katalon, apache Jmeter, postman, soapUI),

* Postman, merupakan http client untuk melakukan testing web service.

* Open API adalah API open source yang dapat digunakan secara bebas.

*** REST API Design Best Practice

** Gunakan kata kerja dibandingkan kata kerja
* GET /books/123
* DELETE /books/123
* POST /books
* PUT /books/123

** Gunakan kata kerja jamak
* GET /cars/123
* POST /cars 

** Gunakan turunan resource untuk menunjukkan relasi / hirarki
* /users
* /users/123
* /users/123/orders
* /userss/123/orders/0001


*** Filtering, Sorting, Paging, and Field Selection
** Filtering
* GET /users?country=USA

** Sorting
* GET /users?sort=birthdate_date:asc

** Paging
* GET /users?limit=100
* GET /users?offset=2

** Field Selection
* GET /users/123?fields=username,email
* GET /users?fields=username,email


*** Handle garis miring dengan baik
* POST: /cars


*** Versioning
** https://uss6.api.mailchimp.com/3.0 (major + minor version indication)
** https://api.stripe.com/v1/ (major version indication only)
** https://developer.github.com/v3/ (major version indication only)


* untuk membuat API kita akan menggunakan library net/http(package untuk membuat server api dan mengconsume api) dan encoding/json (digunakan untuk decode json ke objek struck/map string/array object/ string).

*** Consuming RESTful API
** Consuming RESTful API
* Define Struct
* type StarWarsPeople struct{
* Name   string `json"name"`
* height string `json"height"`
* mass   string `json"mass"`
* films  []string `json"films"`
*}
* func main(){
* response,_ := http.Get("https://swapi.co/api/people/1")
* responseData,_ := ioutil.ReadAll(response.Body)
* defer response.Body.Close()
* var LukeSkywalker StarWarsPeople
* json.Unmarshal(responseData, &LukeSkywalker)
*fmt.Println(LukeSkywalker.Name)
*fmt.Println(LukeSkywalker.height)
*fmt.Println(LukeSkywalker.mass)
*fmt.Println(LukeSkywalker.films[0])


*** Third Party Library Golang
** Echo
** Go Kit
** Logrus
** gorm.io
** cobra

* Echo, High performance, extensible, minimalist Go web framework.

** Advantage using Echo framework :
* Optimized Router
* Middleware
* Data Rendering
* Scalable
* Data Binding

** Echo merupakan framework yang minimalist karena
* tidak ada driver database / ORM
* tidak ada struktur folder, define struktur dengan bebas
* template engine : https://echo.labstack.com/guide/templates

** minimalist tapi dapat diperluas
* menggunakan GORM untuk ORM
* Go JWT extended for authentication

** Setting up echo
* go get github.com/labstack/echo/v4
* note : apabila belum ada go.mod lakukan go mod init github.com/labstack/echo/v4

** Basic Routing & Controller (pada file Belajar\belajar Go)

** Data Rendering (pada file Belajar\belajar Go)

** Retrieve Data (pada file Belajar\belajar Go)

** Binding Data (pada file Belajar\belajar Go)
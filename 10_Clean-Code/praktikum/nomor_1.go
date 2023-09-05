package main

//penulisan code rata kiri

type user struct {

id int

username int //username menggunakan int

password int //password menggunakan int

}


type userservice struct {

t []user //nama variabel tidak jelas

}


func (u userservice) getallusers() []user {

return u.t //mengembalikan nilai dengan variabel yang tidak jelas

}


func (u userservice) getuserbyid(id int) user {

for _, r := range u.t { //nama variabel tidak jelas

if id == r.id { //nama variabel tidak jelas

return r //nama variabel tidak jelas

}

}


return user{}

}

/*
a. 8

b dan c sesuai komen diatas
*/
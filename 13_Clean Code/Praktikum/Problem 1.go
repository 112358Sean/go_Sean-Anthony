package main

type user struct {
	id       int
	username int
	password int
}
//mengubah nama "user" menjadi "User"
//mengubah username dan password menjadi string

type userservice struct {
	t []user
}
//mengubah nama "userservice" menjadi "UserService"
//mengubah nama "t" menjadi "users" agar lebih mudah dipahami

func (u userservice) getallusers() []user {
	return u.t
}
//mengubah nama "getallusers" menjadi "GetAllUsers"

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
//mengubah nama "getuserbyid" menjadi "GetUserById"

//menambahkan func main untuk mengolah hasil login
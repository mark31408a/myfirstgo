package user

type User struct {
	Id        	string	`json:"id"`			//post id
	Name		string	`json:"name"`		//username
	Email		string	`json:"email"`		//
	Description	string 	`json:"description"`	//blogger content

}

type CreateUserDto struct{
	Name		string	`json:"name"`		//username
	Email		string	`json:"email"`		//
	Description	string 	`json:"description"`	//blogger content
}
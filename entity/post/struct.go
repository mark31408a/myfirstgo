package post

type Post struct {
	Id        	int		`json:"id"`			//post id
	UserId		int		`json:"user_id"`	//user id 
	Title		string	`json:"title"`		//post title
	Time		string	`json:"time"`		//
	Content 	string	`json:"content"`	//
}
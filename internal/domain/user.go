package domain

type Users struct{
	ID          	 int    `json:"id"`
	UserName         string    `json:"user_name"`
	Password 		 string `json:"password"`
	Mail     		 string `json:"mail"`
	
}
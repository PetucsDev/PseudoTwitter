package domain

type Likes struct{
	ID          	int    `json:"id"`
	UsuariosId 		int `json:"usuarios_id"`
	PublicacionesId int `json:"publicaciones_id"`
}
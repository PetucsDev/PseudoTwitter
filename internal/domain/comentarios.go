package domain

type Comments struct{
	ID          	int    `json:"id"`
	Descripcion   string    `json:"descripcion"`
	UsuariosId 		int `json:"usuarios_id"`
	PublicacionesId int `json:"publicaciones_id"`
}
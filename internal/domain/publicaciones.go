package domain
type Publications struct{
	ID          	int    `json:"id"`
	Titulo   		string    `json:"titulo"`
	Fecha 			string `json:"fecha"`
	UsuariosId 		int `json:"usuarios_id"`
	
}
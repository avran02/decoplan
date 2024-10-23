export interface IUser {
	id: number
	username: string
	email: string
}

export interface IFormData extends Pick<IUser, 'email'> {
	password: string
}

export interface IFormLogin { 
	username: string
	password: string
}

export interface IFormRegister {
	username: string
	email: string
	password: string
}
import { axiosClassic } from '@/api/axios'
import { IFormLogin, IFormRegister } from '@/types/auth.types'
import { getAccessToken, removeFromStorage, saveTokenStorage } from './auth.helper'

interface IAuthResponse {
	id: string
	accessToken: string
}

export enum EnumTokens {
	'ACCESS_TOKEN' = 'accessToken',
	'REFRESH_TOKEN' = 'refreshToken',
}

class AuthService {
	async login( data: IFormLogin ) {
		const response = await axiosClassic.post<IAuthResponse>(
			`/login`,
			data
		)

		if (response.data.accessToken) saveTokenStorage(response.data.accessToken)
		if (response.data.id) localStorage.setItem('userId', response.data.id)

		return response
	}

	async register( data: IFormRegister ) {
		const response = await axiosClassic.post<IAuthResponse>(
			`/register`,
			data
		)

		if (response.data.accessToken) saveTokenStorage(response.data.accessToken)
		if (response.data.id) localStorage.setItem('userId', response.data.id)

		return response
	}


	async getNewTokens() {
		const response = await axiosClassic.post<IAuthResponse>(
			'/refresh-tokens',
			{},
			{
				withCredentials: true,
			}
		)
		return response.data
	}

	async logout() {
		const accessToken = getAccessToken()
		console.log(accessToken);
		
		const response = await axiosClassic.post<boolean>('/logout', 
			{"accessToken" : accessToken}
		)
		console.log(response)
		

		if (response.data) removeFromStorage()
		console.log(response)
		if (response.data) localStorage.removeItem('userId')


		return response
	}
}

export default new AuthService()

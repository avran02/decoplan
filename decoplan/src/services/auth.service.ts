import { axiosClassic } from '@/api/axios'
import { IFormLogin, IFormRegister } from '@/types/auth.types'
import { getAccessToken, removeFromStorage, saveTokenStorage } from './auth.helper'

interface IAuthResponse {
	accessToken: string
}
interface IRegister {
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

		return response
	}

	async register( data: IFormRegister ) {
		const response = await axiosClassic.post<IRegister>(
			`/register`,
			data
		)

		if (response.data.accessToken) saveTokenStorage(response.data.accessToken)
		
			console.log(response.data)

		return response
	}


	async getNewTokens(refreshToken: string) {
		const response = await axiosClassic.post<IAuthResponse>(
			'/refresh-token',
			refreshToken
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


		return response
	}
}

export default new AuthService()

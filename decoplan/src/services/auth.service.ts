import { axiosClassic } from '@/api/axios'
import { IFormLogin, IFormRegister } from '@/types/auth.types'
import { removeFromStorage, saveTokenStorage } from './auth.helper'

interface IAuthResponse {
	accessToken: string
	refreshToken: string
}
interface IRegister {
	access: boolean
}

export enum EnumTokens {
	'ACCESS_TOKEN' = 'accessToken',
	'REFRESH_TOKEN' = 'refreshToken',
}

class AuthService {
	async login( data: IFormLogin) {
		const response = await axiosClassic.post<IAuthResponse>(
			`/auth/login`,
			data
		)

		if (response.data.accessToken) saveTokenStorage(response.data.accessToken)

		return response
	}

	async register( data: IFormRegister) {
		const response = await axiosClassic.post<IRegister>(
			`/auth/register`,
			data
		)
		return response
	}


	async getNewTokens(refreshToken: string) {
		const response = await axiosClassic.post<IAuthResponse>(
			'/auth/refresh',
			refreshToken
		)
		return response.data
	}

	async logout() {
		const response = await axiosClassic.post<boolean>('/auth/logout')

		if (response.data) removeFromStorage()

		return response
	}
}

export default new AuthService()

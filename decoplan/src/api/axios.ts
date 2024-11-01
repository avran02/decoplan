import { API_URL } from '@/constants/constants'
import { getAccessToken } from '@/services/auth.helper'
import axios, { CreateAxiosDefaults } from 'axios'
import { getContentType } from './api.helper'

const axiosOptions: CreateAxiosDefaults = {
	baseURL: API_URL,
	headers: getContentType(),
	withCredentials: true,
}

export const axiosClassic = axios.create(axiosOptions)

export const instance = axios.create(axiosOptions)

instance.interceptors.request.use(config => {
	const accessToken = getAccessToken()

	if (config?.headers && accessToken)
		config.headers.Authorization = `Bearer ${accessToken}`

	return config
})

// instance.interceptors.response.use(
// 	config => config,
// 	async error => {
// 		const originalRequest = error.config

// 		if (
// 			(error?.response?.status === 401) &&
// 			error.config &&
// 			!error.config._isRetry
// 		) {
// 			originalRequest._isRetry = true
// 			try {
// 				await authService.getNewTokens()
// 				return instance.request(originalRequest)
// 			} catch (error) {
// 				if (
// 					errorCatch(error) === 'Invalid or expired refresh token'
// 				)
// 					removeFromStorage()
// 			}
// 		}

// 		throw error
// 	}
// )

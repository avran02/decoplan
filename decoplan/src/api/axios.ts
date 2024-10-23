import { API_URL } from '@/constants/constants'
import axios, { CreateAxiosDefaults } from 'axios'
import { getContentType } from './api.helper'

const axiosOptions: CreateAxiosDefaults = {
	baseURL: API_URL,
	headers: getContentType(),
	withCredentials: true,
}

export const axiosClassic = axios.create(axiosOptions)

export const instance = axios.create(axiosOptions)

// instance.interceptors.request.use(config => {
// 	const accessToken = getAccessToken()

// 	if (config?.headers && accessToken)
// 		config.headers.Authorization = `Bearer ${accessToken}`

// 	return config
// })

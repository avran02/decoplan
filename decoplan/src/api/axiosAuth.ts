import { AUTH_URL } from '@/constants/constants'
import axios, { CreateAxiosDefaults } from 'axios'
import { getContentType } from './api.helper'

const axiosOptions: CreateAxiosDefaults = {
	baseURL: AUTH_URL,
	headers: getContentType(),
	withCredentials: true,
}

export const instance = axios.create(axiosOptions)
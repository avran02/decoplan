import { getAccessToken } from '@/services/auth.helper'

export function useAuth() {
	const accessToken = getAccessToken() as string

	return {
		isLoggedIn: !!accessToken,
	}
}
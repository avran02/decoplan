import { instance } from '@/api/axiosUsersChats'
import { IChatResponse } from '@/types/chat.types'
import { IUser } from '@/types/usersChats.types'

class UsersChatsService {
	async createUser(data: IUser) {
		const response = await instance.post<IUser>(
			"/users",
			{
				data
			},
		)

		return response
	}
	async getChats() {
		const response = await instance.get<IChatResponse>("/users/chats")

		return response
	}
}
export default new UsersChatsService()
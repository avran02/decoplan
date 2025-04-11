"use client"

import Field from "@/components/ui/field/Field"
import { Loader } from "@/components/ui/loader/Loader"
import { useAuth } from "@/hooks/useAuth"
import { useDebounce } from "@/hooks/useDebounce"
import usersChatsService from "@/services/usersChats.service"
import { useQuery } from "@tanstack/react-query"
import { Search } from "lucide-react"
import { useEffect, useState } from "react"
import { ChatListItem } from "./ChatListItem"

export function ChatsList() {
	const { isLoggedIn } = useAuth()
	const [searchTerm, setSearchTerm] = useState("")
	const debouncedSearchTerm = useDebounce(searchTerm)

	const { data, isLoading, isFetching } = useQuery({
		queryKey: ["chats", debouncedSearchTerm],
		queryFn: () => usersChatsService.getChats(),
		enabled: isLoggedIn,
	})

	const [isClient, setIsClient] = useState(false)
	useEffect(() => {
		setIsClient(true)
	}, [])

	if (!isClient) return null

	return (
		<div>
			<div className='border-t border-b border-border p-layout'>
				<Field
					placeholder='Search chats'
					Icon={Search}
					value={searchTerm}
					onChange={(e) => setSearchTerm(e.target.value)}
				/>
			</div>
			<div>
				{isLoading || isFetching ? (
					<div className='p-layout'>
						<Loader />
					</div>
				) : (data?.data?.chats?.length || 0) > 0 ? (
					data?.data.chats.map((chat) => (
						<ChatListItem key={chat.id} chat={chat} />
					))
				) : (
					<div className='p-layout'>No chats available</div>
				)}
			</div>
		</div>
	)
}

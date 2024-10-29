"use client"
// import { Chat } from "@/components/screens/chat/Chat"
import { getAccessToken } from "@/services/auth.helper"
import authService from "@/services/auth.service"
import { useMutation } from "@tanstack/react-query"
import { LogOut } from "lucide-react"
import { useRouter } from "next/navigation"
import useWebSocket from "react-use-websocket"

export default function ChatPage() {
	const accessToken = getAccessToken() || ""
	const { push } = useRouter()
	const { sendMessage, lastMessage, readyState } = useWebSocket(
		"ws://localhost:57151/ws/orderbook"
	)
	const { mutate: mutateLogout, isPending: isLogoutPending } = useMutation({
		mutationKey: ["logout"],
		mutationFn: () => authService.logout(),
		onSuccess() {
			push("/login")
		},
	})
	return (
		<div>
			{/* <Chat token={accessToken} chatId={"1"} /> */}

			<button onClick={() => mutateLogout()}>
				<LogOut />
			</button>
		</div>
		// <div className='w-8/5 border-r border-border h-full'>
		// 	<MessageField />
		// </div>
	)
}

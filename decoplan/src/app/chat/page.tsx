"use client"
import { Chat } from "@/components/screens/chat/Chat"
import { getAccessToken } from "@/services/auth.helper"
import authService from "@/services/auth.service"
import { useMutation } from "@tanstack/react-query"
import { useRouter } from "next/navigation"

export default function ChatPage() {
	const accessToken = getAccessToken() as string
	const { push } = useRouter()
	const { mutate: mutateLogout, isPending: isLogoutPending } = useMutation({
		mutationKey: ["logout"],
		mutationFn: () => authService.logout(),
		onSuccess() {
			push("/login")
		},
	})
	return (
		<div className='h-screen w-screen flex flex-col'>
			<Chat
				token={accessToken}
				chatId={"6c4d2071-c9c9-4432-9b90-14775b3dc708"}
			/>

			{/* <button onClick={() => mutateLogout()}>
				<LogOut />
			</button> */}
		</div>
	)
}

"use client"
import { Loader } from "@/components/ui/loader/Loader"
import useChatWebSocket from "@/hooks/useChatWebSocket"
import React, { useEffect } from "react"
import { Message } from "./Message"
import { MessageField } from "./MessageField"

interface ChatProps {
	token: string
	chatId: string
}

export const Chat: React.FC<ChatProps> = ({ token, chatId }) => {
	const { connect, messages, sendMessage, fetchMessages, readyState } =
		useChatWebSocket(token)
	const isLoading = false

	useEffect(() => {
		connect()
		fetchMessages({ chatId })
	}, [chatId, fetchMessages])

	const handleSendMessage = (text: string) => {
		const timestamp = new Date().toISOString()
		sendMessage({ chatId, content: { text }, timestamp })
	}

	return (
		<div
			className='h-full grid'
			style={{
				gridTemplateRows: isLoading ? "1fr .089fr" : "1fr .089fr",
			}}
		>
			{isLoading ? (
				<div className='flex items-center justify-center'>
					<Loader />
				</div>
			) : (
				<>
					{/* <ChatHeader correspondent={correspondent} /> */}
					<div className='p-layout border-t border-border'>
						{messages.map((message) => (
							<Message key={message.id} message={message} />
						))}
					</div>
				</>
			)}
			<MessageField onSendMessage={handleSendMessage} />
		</div>
	)
}

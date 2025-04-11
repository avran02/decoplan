"use client"
import { Loader } from "@/components/ui/loader/Loader"
import useChatWebSocket from "@/hooks/useChatWebSocket"
import { getAccessToken } from "@/services/auth.helper"
import { IAttachment, INewMessageDto } from "@/types/chat.types"
import { useEffect, useRef, useState } from "react"
import { Message } from "./Message"
import { MessageField } from "./MessageField"

interface ChatProps {
	chatId: string
}

export function Chat({ chatId }: ChatProps) {
	const { connect, messages, sendMessage, fetchMessages } = useChatWebSocket(
		getAccessToken() as string
	)
	const isLoading = false
	const [attachments, setAttachments] = useState<IAttachment[]>([])

	const messagesEndRef = useRef<HTMLDivElement | null>(null)

	useEffect(() => {
		connect()
		fetchMessages({ chatId, limit: 100, offset: 0 })
	}, [chatId])

	useEffect(() => {
		if (messagesEndRef.current) {
			messagesEndRef.current.scrollIntoView({ behavior: "smooth" })
		}
	}, [messages])

	const handleSendMessage = (text: string) => {
		const timestamp = new Date().toISOString()
		const messageData: INewMessageDto = {
			timestamp,
			chatId,
			content: {
				text,
				attachments,
			},
		}
		sendMessage(messageData)
		setAttachments([])
	}

	const handleAddAttachment = (attachment: IAttachment) => {
		setAttachments((prev) => [...prev, attachment])
	}
	return (
		<div className='h-full grid' style={{ gridTemplateRows: "1fr .089fr" }}>
			{isLoading ? (
				<div className='flex items-center justify-center'>
					<Loader />
				</div>
			) : (
				<>
					{/* <ChatHeader correspondent={correspondent} /> */}
					<div className='p-layout border-t border-border	overflow-y-auto'>
						{messages.map((message) => (
							<Message key={message.id} message={message} />
						))}
						<div ref={messagesEndRef} />
					</div>
				</>
			)}
			<MessageField onSendMessage={handleSendMessage} />
			{/* Optionally add attachment handling component */}
		</div>
	)
}

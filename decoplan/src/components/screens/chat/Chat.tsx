import { Loader } from "@/components/ui/loader/Loader"
import useChatWebSocket from "@/hooks/useChatWebSocket"
import { IAttachment, INewMessageDto } from "@/types/chat.types"
import { useEffect, useState } from "react"
import { Message } from "./Message"
import { MessageField } from "./MessageField"

interface ChatProps {
	token: string
	chatId: string
}

export function Chat({ token, chatId }: ChatProps) {
	const { connect, messages, sendMessage, fetchMessages, readyState } =
		useChatWebSocket(token)
	const isLoading = false
	const [attachments, setAttachments] = useState<IAttachment[]>([])

	useEffect(() => {
		connect()
		fetchMessages({ chatId, limit: 100, offset: 0 })
		console.log(messages)
	}, [chatId, fetchMessages])

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
		console.log(messageData)
		sendMessage(messageData)
		setAttachments([])
	}

	const handleAddAttachment = (attachment: IAttachment) => {
		setAttachments((prev) => [...prev, attachment])
	}

	console.log(messages)
	return (
		<div
			className='h-full grid'
			style={{ gridTemplateRows: isLoading ? "1fr .089fr" : "1fr .089fr" }}
		>
			{isLoading ? (
				<div className='flex items-center justify-center'>
					<Loader />
				</div>
			) : (
				<>
					{/* <ChatHeader correspondent={correspondent} /> */}
					<div className='p-layout border-t border-border'>
						{messages.map((message, i) => (
							<Message key={i} message={message} />
						))}
					</div>
				</>
			)}
			<MessageField onSendMessage={handleSendMessage} />
			{/* Optionally add attachment handling component */}
		</div>
	)
}

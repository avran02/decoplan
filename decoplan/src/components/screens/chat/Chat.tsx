import { Loader } from "@/components/ui/loader/Loader"
import useChatWebSocket from "@/hooks/useChatWebSocket"
import { IAttachment, INewMessageDto } from "@/types/chat.types"
import { useEffect, useRef, useState } from "react"
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

	// Добавляем реф для контейнера сообщений
	const messagesEndRef = useRef<HTMLDivElement | null>(null)

	useEffect(() => {
		connect()
		fetchMessages({ chatId, limit: 100, offset: 0 })
	}, [chatId])

	// Прокручиваем до последнего сообщения при изменении списка сообщений
	useEffect(() => {
		if (messagesEndRef.current) {
			messagesEndRef.current.scrollIntoView({ behavior: "smooth" })
		}
	}, [messages]) // Зависимость на сообщения

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
					<div className='p-layout border-t border-border	overflow-y-auto'>
						{messages.map((message, i) => (
							<Message key={i} message={message} />
						))}
						{/* Добавляем скрытый элемент, чтобы прокручивать к последнему сообщению */}
						<div ref={messagesEndRef} />
					</div>
				</>
			)}
			<MessageField onSendMessage={handleSendMessage} />
			{/* Optionally add attachment handling component */}
		</div>
	)
}

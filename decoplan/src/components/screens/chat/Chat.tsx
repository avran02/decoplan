// "use client"
// import { MessageContent } from "@/types/chat.types"
// import React, { useEffect, useState } from "react"
// import { ReadyState } from "react-use-websocket"

// interface ChatProps {
// 	token: string
// 	chatId: string
// }

// export const Chat: React.FC<ChatProps> = ({ token, chatId }) => {
// 	const [messageText, setMessageText] = useState<string>("")
// 	// const {
// 	// 	messages,
// 	// 	connect,
// 	// 	disconnect,
// 	// 	sendMessage,
// 	// 	fetchMessages,
// 	// 	deleteMessage,
// 	// 	readyState,
// 	// } = useChatWebSocket(token)

// 	// useEffect(() => {
// 	// 	connect()
// 	// 	fetchMessages({ chatId, limit: 10, offset: 0 })
// 	// 	return () => disconnect()
// 	// }, [chatId, connect, disconnect, fetchMessages])

// 	// const handleSendMessage = () => {
// 	// 	const newMessage: MessageContent = {
// 	// 		text: messageText,
// 	// 		attachments: [],
// 	// 	}
// 	// 	sendMessage({
// 	// 		chatId,
// 	// 		content: newMessage,
// 	// 		timestamp: new Date().toISOString(),
// 	// 	})
// 	// 	setMessageText("")
// 	// }

// 	// const handleDeleteMessage = (messageId: string) => {
// 	// 	deleteMessage({ chatId, messageId: Number(messageId) })
// 	// }

// 	return (
// 		<div>
// 			<h2>Чат</h2>
// 			<div className='chat-messages text-white'>
// 				{messages.map((msg) => (
// 					<div key={msg.id} className='message'>
// 						<p>
// 							<strong>{msg.sender}</strong>: {msg.content.text}
// 						</p>
// 						{msg.content.attachments?.map((att) => (
// 							<img
// 								key={att.id}
// 								src={att.url}
// 								alt='attachment'
// 								style={{ width: "100px" }}
// 							/>
// 						))}
// 						<button onClick={() => handleDeleteMessage(msg.id)}>Удалить</button>
// 					</div>
// 				))}
// 			</div>
// 			<input
// 				className='text-black'
// 				type='text'
// 				value={messageText}
// 				onChange={(e) => setMessageText(e.target.value)}
// 				placeholder='Введите сообщение...'
// 			/>
// 			<button
// 				onClick={handleSendMessage}
// 				disabled={readyState !== ReadyState.OPEN}
// 			>
// 				Отправить
// 			</button>
// 		</div>
// 	)
// }

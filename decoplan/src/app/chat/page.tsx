"use client"
import useWebSocket from "react-use-websocket"

export default function ChatPage() {
	const { sendMessage, lastMessage, readyState } = useWebSocket(
		"ws://localhost:57151/ws/orderbook"
	)
	return (
		<div>
			<h1>Chat</h1>
			<div>ReadyState: {readyState}</div>
			<div>
				Last message: {lastMessage ? JSON.stringify(lastMessage?.data) : "None"}
			</div>
			<div>
				<button
					onClick={() => {
						sendMessage("Hello Server!")
					}}
				>
					Send
				</button>
			</div>
		</div>
	)
}

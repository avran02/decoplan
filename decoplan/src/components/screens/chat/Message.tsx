import { IMessage } from "@/types/chat.types"
import dayjs from "dayjs"

export function Message({ message }: { message: IMessage }) {
	const isSender = localStorage.getItem("userId") === message.sender
	return (
		<div
			className={`flex ${isSender ? "justify-end" : "justify-start"} mb-2.5`}
		>
			<div
				className={`relative flex items-center ${
					isSender ? "flex-row-reverse" : ""
				}`}
			>
				<div className={isSender ? "mr-3" : "ml-3"}>
					<div
						className={`text-sm text-white py-1.5 mt-4 px-3 rounded-2xl ${
							isSender
								? "rounded-tr-none bg-primary"
								: "rounded-tl-none bg-border"
						}`}
					>
						{message.content.text}
					</div>
					<div
						className={`text-xs opacity-30 block mt-1.5 ${
							isSender ? "text-right" : "text-left"
						}`}
					>
						{dayjs(message.timestamp).format("HH:mm")}
					</div>
				</div>
			</div>
		</div>
	)
}

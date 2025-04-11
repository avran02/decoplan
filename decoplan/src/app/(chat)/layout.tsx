import { ChatsList } from "@/components/screens/chats/list/ChatsList"
import { type PropsWithChildren } from "react"

export default function ChatLayout({ children }: PropsWithChildren<unknown>) {
	return (
		<div
			className='grid h-full relative'
			style={{
				gridTemplateColumns: ".7fr 3fr",
			}}
		>
			<div className='border-r border-border'>
				<ChatsList />
			</div>
			<div>{children}</div>
		</div>
	)
}

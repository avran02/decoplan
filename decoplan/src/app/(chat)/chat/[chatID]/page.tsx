import type { Metadata } from "next"

import { Chat } from "@/components/screens/chats/chat/Chat"
import { NO_INDEX_PAGE } from "@/constants/seo.constants"

export const metadata: Metadata = {
	...NO_INDEX_PAGE,
}

export default function ChatPage({ params }: { params: { chatID: string } }) {
	return (
		<div className='h-screen flex flex-col'>
			<Chat chatId={params.chatID} />
		</div>
	)
}
